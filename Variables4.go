package main

import (
	"fmt"
)

//<-- Go Multiple Variable Declaration -->//

func main() {

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) // Only value of "d" will be Printed

	// NOTE -> if you use "type" keyword, it is only possible to declare one type of variable per line
}
