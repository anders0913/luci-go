// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/cipd/client/cipd/plugin/protocol"
)

// ErrTerminated is returned by Process.Err() if the plugin process exited with
// 0 exit code.
var ErrTerminated = errors.Reason("terminated with 0 exit code").Err()

// Host launches plugin subprocesses and accepts connections from them.
type Host struct {
	ServiceURL string // URL of the CIPD repository ("https://...") used by the client

	m       sync.Mutex          // protects all fields below
	plugins map[string]*Process // all launched plugins, per their RPC ticket
	srv     *grpc.Server        // non-nil if the gRPC server has started
	srvErr  error               // non-nil if the gRPC server failed to start
	port    int                 // a localhost TCP port the server is listening on

	testServeErr error // non-nil in tests to simulate srv.Serve error
}

// Controller lives in the host process and manages communication with a plugin.
//
// Each instance of a plugin has a Controller associated with it. The controller
// exposes a subset of gRPC services that the particular plugin is allowed to
// call. That way it's possible to have different kinds of plugins by exposing
// different services to them.
type Controller struct {
	Admissions protocol.AdmissionsServer // non-nil for deployment admission plugins
}

// LaunchPlugin launches a plugin subprocesses.
//
// Does not wait for it to connect. Returns a handle that can be used to control
// the process.
//
// Uses the given context for logging from the plugin.
func (h *Host) LaunchPlugin(ctx context.Context, args []string, ctrl *Controller) (*Process, error) {
	if len(args) == 0 {
		return nil, errors.Reason("need at least one argument (the executable to start)").Err()
	}

	// Generate a string sent to the plugin via the stdin. It is used to
	// authenticate the gRPC calls from it.
	ticketBytes := make([]byte, 64)
	if _, err := rand.Read(ticketBytes); err != nil {
		return nil, errors.Annotate(err, "failed to generate random string").Err()
	}
	ticket := base64.RawStdEncoding.EncodeToString(ticketBytes)

	// Initialize the server asynchronously.
	serverReady := h.ensureServerRunning()

	// While the server is being initialized, launch the plugin subprocess.
	plugin, err := h.launchProcess(ctx, args, ctrl, ticket)

	// Wait for the server to start, since we need its port now. Do it even if
	// we failed to start the plugin (which should be rare), to avoid leaking
	// the running goroutine.
	<-serverReady

	// Bail if the plugin failed to start (e.g. a wrong command line).
	if err != nil {
		return nil, err
	}

	// If the server was launched successfully, add the plugin to the plugin set,
	// so Close() terminates it.
	h.m.Lock()
	port := h.port
	if h.srvErr != nil {
		err = errors.Annotate(h.srvErr, "failed to launch the plugins grpc server").Err()
	} else {
		if h.plugins == nil {
			h.plugins = map[string]*Process{}
		}
		h.plugins[ticket] = plugin
	}
	h.m.Unlock()

	// Tell the plugin how to connect to us.
	if err == nil {
		err = plugin.sendHandshake(ctx, port, ticket)
	}

	// If something went wrong, give up and terminate the plugin.
	if err != nil {
		plugin.Terminate(ctx)
		return nil, err
	}

	return plugin, nil
}

// Close terminates all plugins and releases all resources.
//
// Waits for the plugins to terminate gracefully. On context deadline kills them
// with SIGKILL.
func (h *Host) Close(ctx context.Context) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	h.m.Lock()
	defer h.m.Unlock()

	// Note: see also serverCrashed(...).
	if h.srvErr == nil {
		h.srvErr = errors.Reason("the plugins host is already closed").Err()
	}
	if h.srv != nil {
		h.srv.Stop()
		h.srv = nil
	}

	wg.Add(len(h.plugins))
	for _, plugin := range h.plugins {
		plugin := plugin
		go func() {
			defer wg.Done()
			plugin.Terminate(ctx)
		}()
	}
}

// ensureServerRunning launches the localhost server that plugins connect to.
//
// Returns a channel that closes when the server has started successfully or
// fails to start.
func (h *Host) ensureServerRunning() chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		h.m.Lock()
		defer h.m.Unlock()

		if h.srv == nil && h.srvErr == nil {
			var listener net.Listener
			listener, h.srvErr = net.Listen("tcp", "127.0.0.1:0")
			if h.srvErr == nil {
				h.port = listener.(*net.TCPListener).Addr().(*net.TCPAddr).Port
				h.srv = h.initServer(listener)
			}
		}
	}()

	return ch
}

// initServer initializes h.srv and launches the serving loop.
func (h *Host) initServer(listener net.Listener) *grpc.Server {
	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{Time: 5 * time.Second}),
	)

	protocol.RegisterHostServer(srv, &hostServer{host: h})
	protocol.RegisterAdmissionsServer(srv, &admissionsRouter{host: h})

	go func() {
		err := h.testServeErr
		if err == nil {
			err = srv.Serve(listener)
		}
		if err != nil {
			h.serverCrashed(err)
		}
	}()

	return srv
}

// serverCrashed is called from an internal goroutine if the gRPC server
// unexpectedly stops serving.
func (h *Host) serverCrashed(err error) {
	h.m.Lock()
	defer h.m.Unlock()
	if h.srvErr == nil {
		h.srvErr = err
	}
	for _, plugin := range h.plugins {
		go plugin.Terminate(plugin.ctx)
	}
}

// launchProcess starts a plugin subprocess.
//
// Doesn't communicate with it yet. Doesn't update `h.plugins`.
func (h *Host) launchProcess(ctx context.Context, args []string, ctrl *Controller, ticket string) (*Process, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, errors.Annotate(err, "failed to open the stdin pipe").Err()
	}
	if err := cmd.Start(); err != nil {
		return nil, errors.Annotate(err, "failed to launch the plugin").Err()
	}
	proc := &Process{
		ctx:    ctx,
		host:   h,
		ctrl:   ctrl,
		ticket: ticket,
		name:   filepath.Base(args[0]),
		cmd:    cmd,
		stdin:  stdin,
		done:   make(chan struct{}),
	}
	proc.babysit()
	return proc, nil
}

// pluginForRPC picks a *Process based on a ticket in the incoming RPC.
//
// Returns a gRPC status on errors.
func (h *Host) pluginForRPC(ctx context.Context) (*Process, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tickets := md.Get(ticketMetadataKey)
	if len(tickets) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "no ticket in the request")
	}
	h.m.Lock()
	defer h.m.Unlock()
	if plugin, _ := h.plugins[tickets[0]]; plugin != nil {
		return plugin, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "invalid ticket in the request")
}

////////////////////////////////////////////////////////////////////////////////

// Process represents a plugin subprocess.
type Process struct {
	ctx    context.Context // to use for logging from the plugin
	host   *Host           // the owning host
	ctrl   *Controller     // implements the plugin communication logic
	ticket string          // generated random string to use for authentication
	name   string          // identifies the plugin in host logs
	cmd    *exec.Cmd       // the plugin subprocess
	stdin  io.WriteCloser  // a pipe connected to the plugin's stdin
	done   chan struct{}   // closed after the subprocess has terminated
	err    atomic.Value    // an error from cmd.Wait(), valid after "done" closes
}

// Done returns a channel that closes when the plugin terminates.
func (p *Process) Done() <-chan struct{} {
	return p.done
}

// Err returns an error if the plugin terminated (gracefully or not).
//
// Valid only if Done() is already closed, nil otherwise.
func (p *Process) Err() error {
	err, _ := p.err.Load().(error)
	return err
}

// Terminate tries to gracefully terminate the plugin process, killing it with
// SIGKILL on the context expiration or after 5 sec.
//
// Returns the same value as Err(). Does nothing if the process is already
// terminated.
func (p *Process) Terminate(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Make the host "forget" this plugin.
	p.host.m.Lock()
	delete(p.host.plugins, p.ticket)
	p.host.m.Unlock()

	// Ask the plugin to terminate by closing its stdin.
	p.stdin.Close()

	// On context deadline, kill the process forcefully.
	go func() {
		select {
		case <-p.done:
			// Stopped on its own.
		case <-ctx.Done():
			p.cmd.Process.Kill()
		}
	}()

	// Either way wait until the process is reaped.
	<-p.done
	return p.Err()
}

// babysit launches a goroutine that waits for the process to stop.
func (p *Process) babysit() {
	go func() {
		err := p.cmd.Wait()
		if err == nil {
			err = ErrTerminated
		}
		p.err.Store(err)
		close(p.done)
	}()
}

// sendHandshake writes the handshake message to the process's stdin.
//
// Respects the context expiration.
func (p *Process) sendHandshake(ctx context.Context, port int, ticket string) error {
	opts := protojson.MarshalOptions{Multiline: false}
	body, err := opts.Marshal(&protocol.Handshake{Port: int32(port), Ticket: ticket})
	if err != nil {
		return errors.Annotate(err, "failed to serialize the handshake message").Err()
	}
	body = append(body, '\n')

	written := make(chan error, 1)
	go func() {
		_, err := p.stdin.Write(body)
		written <- err
	}()

	select {
	case err := <-written:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-p.done:
		return errors.Reason("the plugin process has unexpectedly terminated").Err()
	}
}

////////////////////////////////////////////////////////////////////////////////

// hostServer implements generally available RPCs.
type hostServer struct {
	protocol.UnimplementedHostServer
	host *Host
}

func (s *hostServer) Log(ctx context.Context, req *protocol.LogRequest) (*emptypb.Empty, error) {
	plugin, err := s.host.pluginForRPC(ctx)
	if err != nil {
		return nil, err
	}

	// Note: `ctx` here is a bare gRPC context without any loggers. Use the plugin
	// context instead.
	lvl := logging.Info
	lvl.Set(req.Severity)
	logging.Logf(plugin.ctx, lvl, "%s: %s", plugin.name, req.Message)

	return &emptypb.Empty{}, nil
}

////////////////////////////////////////////////////////////////////////////////

// admissionsRouter routes Admissions RPCs to the plugin's controller.
type admissionsRouter struct {
	protocol.UnsafeAdmissionsServer
	host *Host
}

func (s *admissionsRouter) impl(ctx context.Context) (protocol.AdmissionsServer, error) {
	switch plugin, err := s.host.pluginForRPC(ctx); {
	case err != nil:
		return nil, err
	case plugin.ctrl.Admissions == nil:
		return nil, status.Errorf(codes.Unimplemented, "not available for this plugin kind")
	default:
		return plugin.ctrl.Admissions, nil
	}
}

func (s *admissionsRouter) ListAdmissions(req *protocol.ListAdmissionsRequest, stream protocol.Admissions_ListAdmissionsServer) error {
	srv, err := s.impl(stream.Context())
	if err != nil {
		return err
	}
	return srv.ListAdmissions(req, stream)
}

func (s *admissionsRouter) ResolveAdmission(ctx context.Context, req *protocol.ResolveAdmissionRequest) (*emptypb.Empty, error) {
	srv, err := s.impl(ctx)
	if err != nil {
		return nil, err
	}
	return srv.ResolveAdmission(ctx, req)
}
