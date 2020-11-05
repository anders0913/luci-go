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

package git

import (
	"go.chromium.org/luci/rts/filegraph"
)

// Graph is a file graph based on the git history.
type Graph struct {
	// Commit is the git commit that the graph state corresponds to.
	Commit string
}

// Node implements filegraph.Graph.
func (g *Graph) Node(name string) filegraph.Node {
	// TODO(crbug.com/1136280): implement.
	panic("not implemented")
}
