// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalComplex128(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"complex128": pfmt.Complex128(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex128":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any complex128": pfmt.Any(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex128":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect complex128": pfmt.Reflect(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex128":"1+23i"
		}`,
		},
	}

	testMarshal(t, tests)
}
