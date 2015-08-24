# Strset [![Build Status](https://travis-ci.org/bsm/strset.png?branch=master)](https://travis-ci.org/bsm/strset)

Simplest-possible string set implementation - uses sorted slices.

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
  set.Add("b") // true
  set.Add("a") // true
  set.Add("c") // true
  set.Add("a") // false

  fmt.Println(set.Slice()) // ["a", "b", "c"]

  set.Exists("a") // true
  set.Exists("d") // false

  set.Remove("a") // true
  set.Remove("d") // false
  fmt.Println(set.Slice()) // ["b", "c"]
}
```

### Licence

```
Copyright (c) 2015 Black Square Media

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
