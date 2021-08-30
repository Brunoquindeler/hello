package main

import (
	"fmt"

	"github.com/Brunoquindeler/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, World!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
