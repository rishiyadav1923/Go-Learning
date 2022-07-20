// Updating & Adding Map Elements
// Synatx -> map_name[key] = value

package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)
	a["Branding"] = "Ford"
	a["Model"] = "Mustang"
	a["Year"] = "1964"

	fmt.Println(a)

	a["Year"] = "1970" //Updating an element
	a["Color"] = "Red" //Adding an element

	fmt.Println(a)
}
