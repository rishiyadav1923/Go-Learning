// Maps are references
package main

import (
	"fmt"
)

func main() {

	var a = map[string]string{"Brand": "Ford", "Model": "Mustang", "Year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["Year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}
