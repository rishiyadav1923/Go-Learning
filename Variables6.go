package main

import (
	"fmt"
)

// <-- Go Variable Declaration in a Block --> //
func main() {
	var (
		a int
		b int    = 1
		c string = "Hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
