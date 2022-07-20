/*
	Allowed Key Types
	-> Booleans
	-> String
	-> Numbers
	-> Arrays
	-> Pointers
	-> Structs
	-> Interfaces (as long as the dynamic type supports equality)

	Invalid Key types are:
	-> Slices
	-> Maps
	-> Functions

	Syntax -> value = map_name[key]
*/

package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)
	a["Brand"] = "Ford"
	a["Mode"] = "Mustang"
	a["Year"] = "1964"

	fmt.Printf(a["Brand"])
}
