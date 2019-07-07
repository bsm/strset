# String Set

[![Build Status](https://travis-ci.org/bsm/strset.png?branch=master)](https://travis-ci.org/bsm/strset)
[![GoDoc](https://godoc.org/github.com/bsm/strset?status.png)](http://godoc.org/github.com/bsm/strset)
[![Go Report Card](https://goreportcard.com/badge/github.com/bsm/strset)](https://goreportcard.com/report/github.com/bsm/strset)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Simplest-possible string set implementation - uses sorted slices.

### Documentation

Full documentation is available on [GoDoc](http://godoc.org/github.com/bsm/strset)

### Example

```go
package main

import (
  "fmt"

  "github.com/bsm/strset"
)

func main() {
	// Create a new set
	set := strset.New(3)
	set.Add("b")	// true
	set.Add("a")	// true
	set.Add("c")	// true
	set.Add("a")	// false

	fmt.Println(set.Slice())	// [a b c]

	set.Has("a")	// true
	set.Has("d")	// false

	set.Remove("a")			// true
	set.Remove("d")			// false
	fmt.Println(set.Slice())	// [b c]

}
```
