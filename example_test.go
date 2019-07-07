package strset_test

import (
	"fmt"

	"github.com/bsm/strset"
)

func ExampleSet() {
	// Create a new set
	set := strset.New(3)
	set.Add("b") // true
	set.Add("a") // true
	set.Add("c") // true
	set.Add("a") // false

	fmt.Println(set.Slice()) // [a b c]

	set.Has("a") // true
	set.Has("d") // false

	set.Remove("a")          // true
	set.Remove("d")          // false
	fmt.Println(set.Slice()) // [b c]

	// Output:
	// [a b c]
	// [b c]
}
