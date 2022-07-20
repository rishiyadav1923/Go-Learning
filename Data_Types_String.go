// String Data Types -> The "string" data type is used to store a sequence of charachters (text).
// String values must be surrounded by double quotes:

package main

import (
	"fmt"
)

func main() {

	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("String Types in Go :- \n")
	fmt.Printf("Type: %T, value: %v \n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v \n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v \n", txt3, txt3)
}
