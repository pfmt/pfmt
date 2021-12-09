// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalTextsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"texts": pfmt.Texts(pfmt.String("Hello, Wörld!"), pfmt.String("Hello, World!"))},
		want:     `Hello, Wörld! Hello, World!`,
		wantText: `Hello, Wörld! Hello, World!`,
		wantJSON: `{
			"texts":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of text jsons": pfmt.Texts(pfmt.String(`{"foo":"bar"}`), pfmt.String("[42]"))},
		want:     `{\"foo\":\"bar\"} [42]`,
		wantText: `{\"foo\":\"bar\"} [42]`,
		wantJSON: `{
			"slice of text jsons":["{\"foo\":\"bar\"}","[42]"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of texts with unescaped null byte": pfmt.Texts(pfmt.String("Hello, Wörld!\x00"), pfmt.String("Hello, World!"))},
		want:     "Hello, Wörld!\\u0000 Hello, World!",
		wantText: "Hello, Wörld!\\u0000 Hello, World!",
		wantJSON: `{
			"slice of texts with unescaped null byte":["Hello, Wörld!\u0000","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty texts": pfmt.Texts(pfmt.String(""), pfmt.String(""))},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of empty texts":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of text nils": pfmt.Texts(nil, nil)},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of text nils":[null,null]
		}`,
	},
}

func TestMarshalTexts(t *testing.T) {
	testMarshal(t, MarshalTextsTests)
}