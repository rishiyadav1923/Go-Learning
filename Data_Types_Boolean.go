// Boolean Data Type
package main

import (
	"fmt"
)

func main() {

	var b1 bool = true // Typed declaration with initial value
	var b2 = true      // Untyped declaration with initial value
	var b3 bool        // Typed declaration without initial value
	b4 := true         // Untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

}
