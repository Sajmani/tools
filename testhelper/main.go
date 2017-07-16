// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The testhelper command converts test helper functions to use t.Helper().
package main

// Usage describes the testhelper command line.
var Usage = `Converts test helper functions to use t.Helper().

  testhelper <function> <args>

testhelper converts a function of the form
  func F(a, b) (x, y, error)
to one of the form:
  func F(t *testing.T, a, b) (x, y)

testhelper adds a call to t.Helper() to the beginning of F's body,
and replaces return statements with non-nil errors with calls to t.Fatal:
  return nil, 0, err
to:
  t.Fatal(err)

testhelper converts calls to F of the form:
  x, y, err := F(a, b)
  if err != nil {
    t.Fatal(err)
  }
to:
  x, y := F(t, a, b)

If F only returned an error, then testhelper
also converts call sites of the form:
  if err := F(a, b); err != nil {
    t.Fatal(err)
  }
to:
  F(t, a, b)
`

func main() {
	
}
