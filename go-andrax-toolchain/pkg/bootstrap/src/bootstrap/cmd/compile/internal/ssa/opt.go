// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/mythical/Downloads/go/src/cmd/compile/internal/ssa/opt.go

//line /home/mythical/Downloads/go/src/cmd/compile/internal/ssa/opt.go:1
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// machine-independent optimization
func opt(f *Func) {
	applyRewrite(f, rewriteBlockgeneric, rewriteValuegeneric)
}
