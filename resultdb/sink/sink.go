// Copyright 2019 The LUCI Authors.
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

// Package sink provides a server for aggregating test results and sending them
// to the ResultDB backend.
package sink

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"go.chromium.org/luci/common/data/rand/cryptorand"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/lucictx"
	"go.chromium.org/luci/server/middleware"
	"go.chromium.org/luci/server/router"

	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

const (
	// DefaultAddr is the TCP address that the Server listens on by default.
	DefaultAddr = "localhost:62115"
	// AuthTokenKey is the key of the HTTP request header where the auth token should
	// be specified. Note that the an HTTP request must have the auth token in the header
	// with the following format.
	// Authorization: ResultSink <auth_token>
	AuthTokenKey = "Authorization"
	// AuthTokenPrefix is embedded into the value of the Authorization HTTP request header,
	// where the auth_token must be present. For the details about the value format of
	// the Authoization HTTP request header, find the description of `AuthTokenKey`.
	AuthTokenPrefix = "ResultSink"
)

// ErrCloseBeforeStart is returned by Close(), when it was invoked before the server
// started.
var ErrCloseBeforeStart error = errors.Reason("the server is not started yet").Err()

// ServerConfig defines the parameters of the server.
type ServerConfig struct {
	// Recorder is the gRPC client to the Recorder service exposed by ResultDB.
	Recorder *pb.RecorderClient

	// AuthToken is a secret token to expect from clients. If it is "" then it
	// will be randomly generated in a secure way.
	AuthToken string
	// Address is the HTTP address to listen on. If empty, the server will use
	// the default address.
	Address string

	// Invocation is the name of the invocation that test results should append
	// to.
	Invocation string
	// UpdateToken is the token that allows writes to Invocation.
	UpdateToken string

	// TestIDPrefix will be prepended to the test_id of each TestResult.
	TestIDPrefix string
}

// Server contains state relevant to the server itself.
// It should always be created by a call to NewServer.
// After a call to Serve(), Server will accept connections on its Port and
// gather test results to send to its Recorder.
type Server struct {
	cfg     ServerConfig
	errC    chan error
	httpSrv http.Server

	// 1 indicates that the server is starting or has started. 0, otherwise.
	started int32

	// Listener for tests
	testListener net.Listener
}

// NewServer creates a Server value and populates optional values with defaults.
func NewServer(cfg ServerConfig) *Server {
	if cfg.Address == "" {
		cfg.Address = DefaultAddr
	}

	s := &Server{
		cfg:  cfg,
		errC: make(chan error, 1),
	}
	return s
}

// Config retrieves the ServerConfig of a previously created Server.
//
// Use this to retrieve the resolved values of unset optional fields in the
// original ServerConfig.
func (s *Server) Config() ServerConfig {
	return s.cfg
}

// ErrC returns a channel that transmits a server error.
//
// Receiving an error on this channel implies that the server has either
// stopped running or is in the process of stopping.
func (s *Server) ErrC() <-chan error {
	return s.errC
}

// Run invokes callback in a context where the server is running.
//
// The context passed to callback will be cancelled if the server encounters
// an error. The context also has the server's information exported into it.
// If callback finishes running, Run will return the error it returned.
func (s *Server) Run(ctx context.Context, callback func(context.Context) error) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// TODO(sajjadm): Add Export here when implemented and explain it in the function comment.
	if err := s.Start(ctx); err != nil {
		return err
	}
	defer s.Close()

	done := make(chan error)
	go func() {
		done <- callback(ctx)
	}()

	select {
	case err := <-s.errC:
		cancel()
		<-done
		return err
	case err := <-done:
		return err
	}
}

// Start runs the server.
//
// On success, Start will return nil, and a subsequent error will be sent on
// the server's ErrC channel.
func (s *Server) Start(ctx context.Context) error {
	if !atomic.CompareAndSwapInt32(&s.started, 0, 1) {
		return errors.Reason("cannot call Start twice").Err()
	}
	if s.cfg.AuthToken == "" {
		t, err := genAuthToken(ctx)
		if err != nil {
			return err
		}
		s.cfg.AuthToken = t
	}

	// launch an HTTP server
	routes := router.NewWithRootContext(ctx)
	routes.Use(router.NewMiddlewareChain(middleware.WithPanicCatcher))
	s.httpSrv.Addr = s.cfg.Address
	s.httpSrv.Handler = routes

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()
		if s.testListener == nil {
			s.errC <- s.httpSrv.ListenAndServe()
		} else {
			// In test mode, the the listener MUST be prepared already.
			s.errC <- s.httpSrv.Serve(s.testListener)
		}
	}()
	go func() {
		<-ctx.Done()
		// TODO(ddoman): handle errors from s.httpSrv.Close()
		s.httpSrv.Close()
	}()
	return nil
}

// Close immediately stops the servers and cancels the requests that are being processed.
func (s *Server) Close() error {
	if atomic.LoadInt32(&s.started) == 0 {
		// hasn't been started
		return ErrCloseBeforeStart
	}
	return s.httpSrv.Close()
}

// Export exports lucictx.ResultSink derived from the server configuration into
// the context.
func (s *Server) Export(ctx context.Context) context.Context {
	return lucictx.SetResultSink(ctx, &lucictx.ResultSink{
		Address:   s.cfg.Address,
		AuthToken: s.cfg.AuthToken,
	})
}

// genAuthToken generates and returns a random auth token.
func genAuthToken(ctx context.Context) (string, error) {
	buf := make([]byte, 32)
	if _, err := cryptorand.Read(ctx, buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// authTokenValue returns the value of the Authorization HTTP header that all requests must
// have.
func authTokenValue(authToken string) string {
	return fmt.Sprintf("%s %s", AuthTokenPrefix, authToken)
}
