// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/mythical/Downloads/go/src/cmd/internal/obj/sort.go

//line /home/mythical/Downloads/go/src/cmd/internal/obj/sort.go:1
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package obj

import "bootstrap/sort"

func SortSlice(slice interface{}, less func(i, j int) bool) {
	sort.Slice(slice, less)
}
