// Boolean Formatting Verbs
// %t -> Value of the Boolean operator in ture or false format (same as using %v)

package main

import (
	"fmt"
)

func main() {

	var i = true
	var j = false

	fmt.Printf("%t\n", i)
	fmt.Printf("%t\n", j)
}
