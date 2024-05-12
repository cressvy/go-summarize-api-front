// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIssue33649(t *testing.T) {
	for _, src := range []string{
		`package p; func _()`,
		`package p; func _() {`,
		`package p; func _() { _ = 0`,
		`package p; func _() { _ = 0 }`,
	} {
		fset := token.NewFileSet()
		f, _ := parser.ParseFile(fset, "", src, parser.AllErrors)
		if f == nil {
			panic("invalid test setup: parser didn't return an AST")
		}

		// find corresponding token.File
		var tf *token.File
		fset.Iterate(func(f *token.File) bool {
			tf = f
			return true
		})
		tfEnd := tf.Base() + tf.Size()

		fd := f.Decls[len(f.Decls)-1].(*ast.FuncDecl)
		fdEnd := int(fd.End())

		if fdEnd != tfEnd {
			t.Errorf("%q: got fdEnd = %d; want %d (base = %d, size = %d)", src, fdEnd, tfEnd, tf.Base(), tf.Size())
		}
	}
}

// TestIssue28089 exercises the IsGenerated function.
func TestIssue28089(t *testing.T) {
	for i, test := range []struct {
		src  string
		want bool
	}{
		// No file comments.
		{`package p`, false},
		// Irrelevant file comments.
		{`// Package p doc.
package p`, false},
		// Special comment misplaced after package decl.
		{`// Package p doc.
package p
// Code generated by gen. DO NOT EDIT.
`, false},
		// Special comment appears inside string literal.
		{`// Package p doc.
package p
const c = "` + "`" + `
// Code generated by gen. DO NOT EDIT.
` + "`" + `
`, false},
		// Special comment appears properly.
		{`// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package p doc comment goes here.
//
// Code generated by gen. DO NOT EDIT.
package p

... `, true},
		// Special comment is indented.
		//
		// Strictly, the indent should cause IsGenerated to
		// yield false, but we cannot detect the indent
		// without either source text or a token.File.
		// In other words, the function signature cannot
		// implement the spec. Let's brush this under the
		// rug since well-formatted code has no indent.
		{`// Package p doc comment goes here.
//
	// Code generated by gen. DO NOT EDIT.
package p

... `, true},
		// Special comment has unwanted spaces after "DO NOT EDIT."
		{`// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package p doc comment goes here.
//
// Code generated by gen. DO NOT EDIT. 
package p

... `, false},
		// Special comment has rogue interior space.
		{`//     Code generated by gen. DO NOT EDIT.
package p
`, false},
		// Special comment lacks the middle portion.
		{`// Code generated DO NOT EDIT.
package p
`, false},
		// Special comment (incl. "//") appears within a /* block */ comment,
		// an obscure corner case of the spec.
		{`/* start of a general comment

// Code generated by tool; DO NOT EDIT.

end of a general comment */

// +build !dev

// Package comment.
package p

// Does match even though it's inside general comment (/*-style).
`, true},
	} {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", test.src, parser.PackageClauseOnly|parser.ParseComments)
		if f == nil {
			t.Fatalf("parse %d failed to return AST: %v", i, err)
		}

		got := ast.IsGenerated(f)
		if got != test.want {
			t.Errorf("%d: IsGenerated on <<%s>> returned %t", i, test.src, got)
		}
	}

}