// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalComplex64s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of complex64s": pfmt.Complex64s([]complex64{complex(1, 23), complex(3, 21)})},
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of complex64s":["1+23i","3+21i"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without complex64s": pfmt.Complex64s(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without complex64s":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c, c2 complex64 = complex(1, 23), complex(3, 21)
				return map[string]json.Marshaler{"slice of any complex64s": pfmt.Anys([]interface{}{c, c2})}
			}(),
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of any complex64s":["1+23i","3+21i"]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c, c2 complex64 = complex(1, 23), complex(3, 21)
				return map[string]json.Marshaler{"slice of reflect complex64": pfmt.Anys([]interface{}{c, c2})}
			}(),
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of reflect complex64":["1+23i","3+21i"]
		}`,
		},
	}

	testMarshal(t, tests)
}
