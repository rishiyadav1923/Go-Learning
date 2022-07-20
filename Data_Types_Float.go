// Float Data-Types
// Type -> Size -> Range
// float32 -> 32-bits -> -3.4e+38 to 3.4e+38
// float64 -> 64-bits -> -1.7e+308 to 1.7e+308
// NOTE -> Default will always be float64 if you not specify

package main

import (
	"fmt"
)

func main() {

	// for float32
	var x float32 = 123.78
	var y float32 = 3.4e+38

	fmt.Printf("For float32 :- \n")
	fmt.Printf("Type: %T, value: %v \n", x, x)
	fmt.Printf("Type: %T, value: %v \n", y, y)

	// for float64
	var p float64 = 123.78
	var q float64 = 1.7e+308

	fmt.Printf("For float64 :- \n")
	fmt.Printf("Type: %T, value: %v \n", p, p)
	fmt.Printf("Type: %T, value: %v \n", q, q)

}
