/*
	Syntax ->
	func FunctionName(parameter-1 type, parameter-2 type, parameter-3 type) {
		code to be executed
	}
*/

package main

import (
	"fmt"
)

// When you are working with Multiple Paramters , the function call must have the same number of arguments as there are parameters,
// and the arguments must be passed in the same order

func family(fname string, age int) {
	fmt.Println("Hello", age, "years-old", fname, "Reference")
}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Reference")
}

func main() {

	familyName("Rishi")
	familyName("Sidharth")
	familyName("Krishna")

	family("Rishi", 3)
	family("Sidharth", 14)
	family("Krishna", 30)

	// When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter
	// while "Rishi", "Sidharth", "Krishna" are arguments

	// Multiple - Parameters

}
