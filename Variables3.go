package main

import (
	"fmt"
)

// declaring & initializing value from outside as well as inside of the Function
var a int
var b int = 2
var c = 3

func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// declaring value inside of the program using ":=" symbol
	// Note -> You can't use ":=" symbol outside the function

	p := 7
	fmt.Println(p)
}
