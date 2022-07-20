// Typed constant V/S Untyped Constants
package main

import (
	"fmt"
)

// Typed Constants -> Typed Constants are declared with a defined type:
const A int = 1

// Untyped Constants -> Untyped Constants are declared without defined type:
const B = 2

func main() {
	fmt.Println(A)
	fmt.Println(B)
}
