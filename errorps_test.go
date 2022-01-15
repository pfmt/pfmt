// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalErrorps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
				return map[string]json.Marshaler{"error pointer slice": pfmt.Errorps([]*error{&err, &err2})}
			}(),
			want:     "something went wrong we have a problem",
			wantText: "something went wrong we have a problem",
			wantJSON: `{
			"error pointer slice":["something went wrong","we have a problem"]
		}`,
		},
		{
			line:  line(),
			input: map[string]json.Marshaler{"nil error pointers": pfmt.Errorps([]*error{nil, nil})},
			wantJSON: `{
			"nil error pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without error pointers": pfmt.Errorps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"without error pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
				return map[string]json.Marshaler{"slice of any error pointers": pfmt.Anys([]interface{}{&err, &err2})}
			}(),
			want:     "something went wrong we have a problem",
			wantText: "something went wrong we have a problem",
			wantJSON: `{
			"slice of any error pointers":["something went wrong","we have a problem"]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
				return map[string]json.Marshaler{"slice of reflect of error pointers": pfmt.Reflects([]interface{}{&err, &err2})}
			}(),
			want:     "{something went wrong} {we have a problem}",
			wantText: "{something went wrong} {we have a problem}",
			wantJSON: `{
			"slice of reflect of error pointers":[{},{}]
		}`,
		},
	}

	testMarshal(t, tests)
}
