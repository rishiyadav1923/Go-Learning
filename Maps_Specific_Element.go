// Checking for aspecific elements in a Map
// Syntax -> OK := map_name[key]

// If you want to check the existence of a certain key, you can use the blank identifier (_) in Place of val

package main

import (
	"fmt"
)

func main() {

	var a = map[string]string{"Brand": "Ford", "Model": "Mustang", "Year": "1964", "day": ""}

	val1, OK1 := a["Brand"] // Checking for existing key and its value
	val2, OK2 := a["Color"] // Checking for existing key and its value
	val3, OK3 := a["Year"]   // Checking for existing key and its value
	_, OK4 := a["Model"]    // Checking for existing key and its value

	fmt.Println(val1, OK1)
	fmt.Println(val2, OK2)
	fmt.Println(val3, OK3)
	fmt.Println(OK4)
}
