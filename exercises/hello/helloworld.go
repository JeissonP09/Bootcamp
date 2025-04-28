package main

import (
	"exercises/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleeH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
