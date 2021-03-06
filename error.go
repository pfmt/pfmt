// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/pfmt/pfmt/pencode"
)

// Err returns stringer/JSON/text marshaler for the error type.
func Err(v error) ErrorV { return New().Err(v) }

// Error returns stringer/JSON/text marshaler for the error type.
func (pretty Pretty) Err(v error) ErrorV {
	return ErrorV{
		v:        v,
		prettier: pretty,
	}
}

type ErrorV struct {
	v        error
	prettier Pretty
}

func (v ErrorV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v ErrorV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	var buf bytes.Buffer

	err := pencode.String(&buf, v.v.Error())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v ErrorV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
