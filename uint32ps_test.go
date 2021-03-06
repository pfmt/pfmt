// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalUint32ps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint32 = 42, 77
				return map[string]json.Marshaler{"uint32 pointer slice": pfmt.Uint32ps([]*uint32{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint32 pointer slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil uint32 pointers": pfmt.Uint32ps([]*uint32{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil uint32 pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint32 pointers": pfmt.Uint32ps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without uint32 pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint32 = 42, 77
				return map[string]json.Marshaler{"slice of any uint32 pointers": pfmt.Anys([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint32 pointers":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint32 = 42, 77
				return map[string]json.Marshaler{"slice of reflects of uint32 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of reflects of uint32 pointers":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
