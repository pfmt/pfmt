// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"fmt"
)

// Float32 returns stringer/JSON/text marshaler for the float32 type.
func Float32(v float32) Float32V { return Float32V{V: v} }

type Float32V struct{ V float32 }

func (v Float32V) String() string {
	return fmt.Sprint(v.V)
}

func (v Float32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Float32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
