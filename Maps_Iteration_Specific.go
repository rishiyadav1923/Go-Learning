// Iterating over Maps in a Specific Order

package main

import (
	"fmt"
)

func main() {

	a := map[string]int{"One": 1, "Two": 2, "Three": 3, "Four": 4}

	var b []string
	b = append(b, "One", "Two", "Three", "Four")

	for k, v := range a {
		// Loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b {
		// Loop with the defined order
		fmt.Printf("%v : %v,", element, a[element])
	}
}
