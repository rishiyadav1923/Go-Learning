// Removing elements is done using the "delete()" function
// Syntax -> delete(map_name, key)

package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)
	a["Brand"] = "Ford"
	a["Model"] = "Mustang"
	a["Year"] = "1954"

	fmt.Println(a)

	delete(a, "Year")

	fmt.Println(a)
}
