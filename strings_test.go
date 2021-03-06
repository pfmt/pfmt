// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalStrings(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"strings": pfmt.Strings([]string{"Hello, Wörld!", "Hello, World!"})},
			want:     "Hello, Wörld! Hello, World!",
			wantText: "Hello, Wörld! Hello, World!",
			wantJSON: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty strings": pfmt.Strings([]string{"", ""})},
			want:     " ",
			wantText: " ",
			wantJSON: `{
			"empty strings":["",""]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"strings with zero byte": pfmt.Strings([]string{string(byte(0)), string(byte(0))})},
			want:     "\\u0000 \\u0000",
			wantText: "\\u0000 \\u0000",
			wantJSON: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without strings": pfmt.Strings(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"without strings":null
		}`,
		},
	}

	testMarshal(t, tests)
}
