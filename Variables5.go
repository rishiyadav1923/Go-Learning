package main

import (
	"fmt"
)

// <-- If the "type" keyword id not specified, you can declaration different types of variables in the same line: --> //
func main() {

	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
