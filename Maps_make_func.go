// Making map using "make()" Function
/*
	Syntax ->
	var a = make(map[keyType]ValueType)
					OR
	b := make(map[keyType]ValueType)
*/

package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)
	// The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["Year"] = "1964"

	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
