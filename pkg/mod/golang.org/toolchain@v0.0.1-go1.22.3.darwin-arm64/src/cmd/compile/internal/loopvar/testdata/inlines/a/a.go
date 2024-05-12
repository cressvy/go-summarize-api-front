// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package a

import "cmd/compile/internal/loopvar/testdata/inlines/b"

func F() []*int {
	var s []*int
	for i := 0; i < 10; i++ {
		s = append(s, &i)
	}
	return s
}

func Fb() []*int {
	bf, _ := b.F()
	return bf
}
