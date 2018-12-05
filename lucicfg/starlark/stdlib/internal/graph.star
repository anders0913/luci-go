# Copyright 2018 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


def _key(*args):
  """Returns a key with given [(kind, name)] path.

  Args:
    *args: even number of strings: kind1, name1, kind2, name2, ...

  Returns:
    graph.key object representing this path.
  """
  return __native__.graph().key(*args)


def _add_node(key, props=None, trace=None):
  """Adds a node to the graph or fails if such node already exists.

  Also fails if used from a generator callback: at this point the graph is
  frozen and can't be extended.

  Args:
    key: a node key, as returned by graph.key(...).
    props: a dict with node properties, will be frozen.
    trace: a stack trace to associate with the node.

  Returns:
    graph.node object representing the node.
  """
  return __native__.graph().add_node(
      key, props or {}, trace or stacktrace(skip=1))


def _add_edge(parent, child, title=None, trace=None):
  """Adds an edge to the graph.

  Neither of the nodes have to exist yet: it is OK to declare nodes and edges
  in arbitrary order as long as at the end of the script execution (when the
  graph is finalized) the graph is complete.

  Fails if there's already an edge between the given nodes with exact same title
  or the new edge introduces a cycle.

  Also fails if used from a generator callback: at this point the graph is
  frozen and can't be extended.

  Args:
    parent: a parent node key, as returned by graph.key(...).
    child: a child node key, as returned by graph.key(...).
    title: a title for the edge, used in error messages.
    trace: a stack trace to associate with the edge.
  """
  return __native__.graph().add_edge(
      parent, child, title or '', trace or stacktrace(skip=1))


def _node(key):
  """Returns a node by the key or None if there's no such node.

  Fails if called not from a generator callback: a graph under construction
  can't be queried.

  Args:
    key: a node key, as returned by graph.key(...).

  Returns:
    graph.node object representing the node.
  """
  return __native__.graph().node(key)


# Public API of this module.
graph = struct(
    key = _key,
    add_node = _add_node,
    add_edge = _add_edge,
    node = _node,
)
