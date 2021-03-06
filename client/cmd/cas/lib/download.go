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

package lib

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/bazelbuild/remote-apis-sdks/go/pkg/client"
	"github.com/bazelbuild/remote-apis-sdks/go/pkg/digest"
	repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/maruel/subcommands"
	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/data/caching/cache"
	"go.chromium.org/luci/common/data/embeddedkvs"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/isolated"
	isol "go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/system/filesystem"
	"go.chromium.org/luci/common/system/signals"
)

const smallFileThreshold = 128 * 1024 // 128KiB

// CmdDownload returns an object for the `download` subcommand.
func CmdDownload(defaultAuthOpts auth.Options) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "download <options>...",
		ShortDesc: "download directory tree from a CAS server.",
		LongDesc: `Downloads directory tree from the CAS server.

Tree is referenced by their digest "<digest hash>/<size bytes>"`,
		CommandRun: func() subcommands.CommandRun {
			c := downloadRun{}
			c.Init(defaultAuthOpts)
			c.cachePolicies.AddFlags(&c.Flags)
			c.Flags.StringVar(&c.cacheDir, "cache-dir", "", "Cache directory to store downloaded files.")
			c.Flags.StringVar(&c.digest, "digest", "", `Digest of root directory proto "<digest hash>/<size bytes>".`)
			c.Flags.StringVar(&c.dir, "dir", "", "Directory to download tree.")
			c.Flags.StringVar(&c.dumpStatsJSON, "dump-stats-json", "", "Dump download stats to json file.")
			c.Flags.StringVar(&c.kvs, "kvs-file", "", "Cache file for small files.")
			return &c
		},
	}
}

type downloadRun struct {
	commonFlags
	digest        string
	dir           string
	dumpStatsJSON string

	cacheDir      string
	cachePolicies cache.Policies

	kvs string
}

func (r *downloadRun) parse(a subcommands.Application, args []string) error {
	if err := r.commonFlags.Parse(); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.Reason("position arguments not expected").Err()
	}

	if r.cacheDir == "" && !r.cachePolicies.IsDefault() {
		return errors.New("cache-dir is necessary when cache-max-size, cache-max-items or cache-min-free-space are specified")
	}

	if r.kvs != "" && r.cacheDir == "" {
		return errors.New("if small-files-cache is set, cache-dir should be set")
	}

	r.dir = filepath.Clean(r.dir)

	return nil
}

func createDirectories(ctx context.Context, root string, outputs map[string]*client.TreeOutput) error {
	logger := logging.Get(ctx)

	start := time.Now()

	dirset := make(map[string]struct{})

	// Extract unique directory paths for optimization.
	for path, output := range outputs {
		var dir string
		if output.IsEmptyDirectory {
			dir = path
		} else {
			dir = filepath.Dir(path)
		}

		for dir != root {
			if _, ok := dirset[dir]; ok {
				break
			}
			dirset[dir] = struct{}{}
			dir = filepath.Dir(dir)
		}
	}

	dirs := make([]string, 0, len(dirset))
	for dir := range dirset {
		dirs = append(dirs, dir)
	}

	sort.Strings(dirs)

	logger.Infof("preprocess took %s", time.Since(start))
	start = time.Now()

	if err := os.MkdirAll(root, 0o700); err != nil {
		return errors.Annotate(err, "failed to create root dir").Err()
	}

	for _, dir := range dirs {
		if err := os.Mkdir(dir, 0o700); err != nil && !os.IsExist(err) {
			return errors.Annotate(err, "failed to create directory").Err()
		}
	}

	logger.Infof("dir creation took %s", time.Since(start))

	return nil
}

func copyFiles(ctx context.Context, dsts []*client.TreeOutput, srcs map[digest.Digest]*client.TreeOutput) error {
	eg, _ := errgroup.WithContext(ctx)

	// limit the number of concurrent I/O operations.
	ch := make(chan struct{}, runtime.NumCPU())

	for _, dst := range dsts {
		dst := dst
		src := srcs[dst.Digest]
		ch <- struct{}{}
		eg.Go(func() (err error) {
			defer func() { <-ch }()
			mode := 0o600
			if dst.IsExecutable {
				mode = 0o700
			}

			if err := filesystem.Copy(dst.Path, src.Path, os.FileMode(mode)); err != nil {
				return errors.Annotate(err, "failed to copy file from '%s' to '%s'", src.Path, dst.Path).Err()
			}

			return nil
		})
	}

	return eg.Wait()
}

func copySmallFilesFromCache(kvs *embeddedkvs.KVS, smallFiles map[string][]*client.TreeOutput) error {
	smallFileHashes := make([]string, 0, len(smallFiles))
	for smallFile := range smallFiles {
		smallFileHashes = append(smallFileHashes, smallFile)
	}

	var mu sync.Mutex

	// limit the number of concurrent I/O operations.
	ch := make(chan struct{}, runtime.NumCPU())

	// Sort hashes by one of corresponding file path.
	sort.Slice(smallFileHashes, func(i, j int) bool {
		filei := smallFiles[smallFileHashes[i]][0]
		filej := smallFiles[smallFileHashes[j]][0]
		return filei.Path < filej.Path
	})

	// Extract small files from kvs.
	return kvs.GetMulti(smallFileHashes, func(key string, value []byte) error {
		ch <- struct{}{}
		defer func() { <-ch }()

		mu.Lock()
		files := smallFiles[key]
		delete(smallFiles, key)
		mu.Unlock()

		for _, file := range files {
			mode := 0o600
			if file.IsExecutable {
				mode = 0o700
			}
			if err := ioutil.WriteFile(file.Path, value, os.FileMode(mode)); err != nil {
				return errors.Annotate(err, "failed to write file").Err()
			}
		}

		return nil
	})
}

func cacheSmallFiles(kvs *embeddedkvs.KVS, outputs []*client.TreeOutput) error {
	var eg errgroup.Group

	bufferPool := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	// limit the number of concurrent I/O operations.
	ch := make(chan struct{}, runtime.NumCPU())

	for _, output := range outputs {
		output := output

		eg.Go(func() error {
			buf := bufferPool.Get().(*bytes.Buffer)
			buf.Reset()
			defer bufferPool.Put(buf)

			b, err := func() ([]byte, error) {
				ch <- struct{}{}
				defer func() { <-ch }()

				f, err := os.Open(output.Path)
				if err != nil {
					return nil, errors.Annotate(err, "failed to open file").Err()
				}
				defer f.Close()

				if _, err := io.Copy(buf, f); err != nil {
					return nil, errors.Annotate(err, "failed to read file").Err()
				}

				return buf.Bytes(), nil
			}()

			if err != nil {
				return err
			}
			return kvs.Set(output.Digest.Hash, b)
		})
	}

	return eg.Wait()
}

func cacheOutputFiles(diskcache *cache.Cache, kvs *embeddedkvs.KVS, outputs map[digest.Digest]*client.TreeOutput) error {
	var smallOutputs, largeOutputs []*client.TreeOutput

	for _, output := range outputs {
		if kvs != nil && output.Digest.Size <= smallFileThreshold {
			smallOutputs = append(smallOutputs, output)
		} else {
			largeOutputs = append(largeOutputs, output)
		}
	}

	// This is to utilize locality of disk access.
	sort.Slice(smallOutputs, func(i, j int) bool {
		return smallOutputs[i].Path < smallOutputs[j].Path
	})

	sort.Slice(largeOutputs, func(i, j int) bool {
		return largeOutputs[i].Path < largeOutputs[j].Path
	})

	if kvs != nil {
		if err := cacheSmallFiles(kvs, smallOutputs); err != nil {
			return err
		}
	}

	for _, output := range largeOutputs {
		if err := diskcache.AddFileWithoutValidation(
			isolated.HexDigest(output.Digest.Hash), output.Path); err != nil {
			return errors.Annotate(err, "failed to add cache; path=%s digest=%s", output.Path, output.Digest).Err()
		}
	}

	return nil
}

// doDownload downloads directory tree from the CAS server.
func (r *downloadRun) doDownload(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	signals.HandleInterrupt(cancel)

	d, err := digest.NewFromString(r.digest)
	if err != nil {
		return errors.Annotate(err, "failed to parse digest: %s", r.digest).Err()
	}

	c, err := newCasClient(ctx, r.casFlags.Instance, r.parsedAuthOpts, true)
	if err != nil {
		return err
	}

	rootDir := &repb.Directory{}
	if err := c.ReadProto(ctx, d, rootDir); err != nil {
		return errors.Annotate(err, "failed to read root directory proto").Err()
	}

	start := time.Now()
	dirs, err := c.GetDirectoryTree(ctx, d.ToProto())
	if err != nil {
		return errors.Annotate(err, "failed to call GetDirectoryTree").Err()
	}
	logger := logging.Get(ctx)
	logger.Infof("finished GetDirectoryTree api call: %d, took %s", len(dirs), time.Since(start))

	start = time.Now()
	t := &repb.Tree{
		Root:     rootDir,
		Children: dirs,
	}

	outputs, err := c.FlattenTree(t, r.dir)
	if err != nil {
		return errors.Annotate(err, "failed to call FlattenTree").Err()
	}

	to := make(map[digest.Digest]*client.TreeOutput)

	var diskcache *cache.Cache
	if r.cacheDir != "" {
		diskcache, err = cache.New(r.cachePolicies, r.cacheDir, crypto.SHA256)
		if err != nil {
			return errors.Annotate(err, "failed to create initialize cache").Err()
		}
		defer diskcache.Close()
	}

	var kvs *embeddedkvs.KVS
	if r.kvs != "" {
		kvs, err = embeddedkvs.New(r.kvs)
		if err != nil {
			return err
		}
		defer kvs.Close()
	}

	if err := createDirectories(ctx, r.dir, outputs); err != nil {
		return err
	}
	logger.Infof("finish createDirectories, took %s", time.Since(start))
	start = time.Now()

	// Files have the same digest are downloaded only once, so we need to
	// copy duplicates files later.
	var dups []*client.TreeOutput

	smallFiles := make(map[string][]*client.TreeOutput)

	for path, output := range outputs {
		if output.IsEmptyDirectory {
			continue
		}

		if output.SymlinkTarget != "" {
			if err := os.Symlink(output.SymlinkTarget, path); err != nil {
				return errors.Annotate(err, "failed to create symlink").Err()
			}
			continue
		}

		if kvs != nil && output.Digest.Size <= smallFileThreshold {
			smallFiles[output.Digest.Hash] = append(smallFiles[output.Digest.Hash], output)
			continue
		}

		if diskcache != nil && diskcache.Touch(isolated.HexDigest(output.Digest.Hash)) {
			mode := 0o600
			if output.IsExecutable {
				mode = 0o700
			}

			if err := diskcache.Hardlink(isolated.HexDigest(output.Digest.Hash), path, os.FileMode(mode)); err != nil {
				return err
			}
			continue
		}

		if _, ok := to[output.Digest]; ok {
			dups = append(dups, output)
		} else {
			to[output.Digest] = output
		}
	}
	logger.Infof("finished copy from cache (if any), dups: %d, to: %d, took %s", len(dups), len(to), time.Since(start))

	start = time.Now()

	if kvs != nil {
		if err := copySmallFilesFromCache(kvs, smallFiles); err != nil {
			return err
		}
	}

	// Process non-cached files.
	for _, files := range smallFiles {
		for _, file := range files {
			if _, ok := to[file.Digest]; ok {
				dups = append(dups, file)
			} else {
				to[file.Digest] = file
			}
		}
	}

	logger.Infof("finished copy small files from cache (if any), to: %d, took %s", len(to), time.Since(start))

	start = time.Now()
	if err := c.DownloadFiles(ctx, "", to); err != nil {
		return errors.Annotate(err, "failed to download files").Err()
	}
	logger.Infof("finished DownloadFiles api call, took %s", time.Since(start))

	if diskcache != nil {
		start = time.Now()
		if err := cacheOutputFiles(diskcache, kvs, to); err != nil {
			return err
		}
		logger.Infof("finished cache addition, took %s", time.Since(start))
	}

	start = time.Now()
	if err := copyFiles(ctx, dups, to); err != nil {
		return err
	}
	logger.Infof("finished files copy of %d, took %s", len(dups), time.Since(start))

	if dsj := r.dumpStatsJSON; dsj != "" {
		cold := make([]int64, 0, len(to))
		for d := range to {
			cold = append(cold, d.Size)
		}
		hot := make([]int64, 0, len(outputs)-len(to))
		for _, output := range outputs {
			d := output.Digest
			if _, ok := to[d]; !ok {
				hot = append(hot, d.Size)
			}
		}

		if err := isol.WriteStats(dsj, hot, cold); err != nil {
			return errors.Annotate(err, "failed to write stats json").Err()
		}
	}

	return nil
}

func (r *downloadRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	ctx := cli.GetContext(a, r, env)
	logging.Get(ctx).Infof("start command")
	if err := r.parse(a, args); err != nil {
		errors.Log(ctx, err)
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	defer r.profiler.Stop()

	if err := r.doDownload(ctx); err != nil {
		errors.Log(ctx, err)
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}

	return 0
}
