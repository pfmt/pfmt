// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalUintptrps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uintptr = 42, 77
				return map[string]json.Marshaler{"uintptr pointer slice": pfmt.Uintptrps([]*uintptr{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uintptr pointer slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil uintptr pointers": pfmt.Uintptrps([]*uintptr{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil uintptr pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uintptr pointers": pfmt.Uintptrps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without uintptr pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uintptr = 42, 77
				return map[string]json.Marshaler{"slice of any uintptr pointers": pfmt.Anys([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uintptr pointers":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uintptr = 42, 77
				return map[string]json.Marshaler{"slice of reflects of uintptr pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of reflects of uintptr pointers":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
