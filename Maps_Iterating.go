// You can use "range" to iterate over maps

package main

import (
	"fmt"
)

func main() {

	a := map[string]int{"One": 1, "Two": 2, "Three": 3, "Four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}
