// Copyright 2018 The LUCI Authors.
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

package builtins

import (
	"encoding/json"
	"fmt"

	"go.starlark.net/starlark"
)

// ToJSON is to_json(value) builtin.
//
//  def to_json(value):
//    """Serializes a value to compact JSON.
//
//    Doesn't support integers that do not fit int64. Fails if the value being
//    converted has cycles.
//
//    Args:
//      value: a starlark value: scalars, lists, tuples, dicts containing only
//        starlark values.
//
//    Returns:
//      A string with its compact JSON serialization.
//    """
var ToJSON = starlark.NewBuiltin("to_json", func(_ *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var v starlark.Value
	if err := starlark.UnpackPositionalArgs(fn.Name(), args, kwargs, 1, &v); err != nil {
		return nil, err
	}
	obj, err := ToGoNative(v)
	if err != nil {
		return nil, err
	}
	blob, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return starlark.String(blob), nil
})

// visitingSet is a set of containers we currently have recursed into.
//
// Used to detect cycles. Note that it is not a stack because we prefer O(1)
// lookup and we know there can't be duplicates in it (so 'remove' is not
// ambiguous).
type visitingSet map[interface{}]struct{}

func (v visitingSet) add(container interface{}) error {
	if _, haveIt := v[container]; haveIt {
		return fmt.Errorf("to_json: detected recursion in the data structure")
	}
	v[container] = struct{}{}
	return nil
}

func (v visitingSet) remove(container interface{}) {
	delete(v, container)
}

// ToGoNative takes a Starlark value and returns native Go value for it.
//
// E.g. it takes *starlark.Dict and returns map[string]interface{}. Works
// recursively.
func ToGoNative(v starlark.Value) (interface{}, error) {
	return toGoNative(v, visitingSet{})
}

// toGoNative implements ToGoNative.
//
// Uses given 'visiting' set to detect cycles in the value being converted, to
// avoid stack overflows due to unbounded recursion.
func toGoNative(v starlark.Value, visiting visitingSet) (interface{}, error) {
	// Add containers to 'visiting' set right away. Note that Tuples are special,
	// since they are not hashable (being a slice). We add a pointer instead.
	var container interface{}
	switch val := v.(type) {
	case starlark.Tuple:
		container = &val
	case *starlark.List, *starlark.Dict, *starlark.Set:
		container = val
	}
	if container != nil {
		if err := visiting.add(container); err != nil {
			return nil, err
		}
		defer visiting.remove(container)
	}

	switch val := v.(type) {
	case starlark.NoneType:
		return nil, nil
	case starlark.Bool:
		return bool(val), nil
	case starlark.Float:
		return float64(val), nil
	case starlark.String:
		return string(val), nil
	case starlark.Int:
		i, ok := val.Int64()
		if !ok {
			return nil, fmt.Errorf("to_json: can't convert %q to int64", val.String())
		}
		return i, nil
	case *starlark.Dict:
		pairs := val.Items()
		out := make(map[string]interface{}, len(pairs))
		for _, pair := range pairs {
			if len(pair) != 2 {
				panic("impossible")
			}
			key, ok := pair[0].(starlark.String)
			if !ok {
				return nil, fmt.Errorf("to_json: dict keys should be strings, got %s", pair[0].Type())
			}
			val, err := toGoNative(pair[1], visiting)
			if err != nil {
				return nil, err
			}
			out[string(key)] = val
		}
		return out, nil
	}

	// This covers *List, Tuple and *Set.
	if iterable, ok := v.(starlark.Iterable); ok {
		iter := iterable.Iterate()
		defer iter.Done()
		out := []interface{}{}
		var val starlark.Value
		for iter.Next(&val) {
			native, err := toGoNative(val, visiting)
			if err != nil {
				return nil, err
			}
			out = append(out, native)
		}
		return out, nil
	}

	return nil, fmt.Errorf("to_json: unsupported type %s", v.Type())
}
