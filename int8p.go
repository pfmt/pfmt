// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int8p returns stringer/JSON/text marshaler for the int8 pointer type.
func Int8p(p *int8) Int8P { return New().Int8p(p) }

// Int8p returns stringer/JSON/text marshaler for the int8 pointer type.
func (pretty Pretty) Int8p(p *int8) Int8P {
	return Int8P{
		p:        p,
		prettier: pretty,
	}
}

type Int8P struct {
	p        *int8
	prettier Pretty
}

func (p Int8P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Int8(*p.p).String()
}

func (p Int8P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int8P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
