package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/mortalpuppet23/hello/morestrings"
)

func main() {
	// fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
