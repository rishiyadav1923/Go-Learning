// In this program we will Learn 3 types of "Print" in Go :-
// Print()
// Println()
// Printf()

package main

import (
	"fmt"
)

func main() {

	// The "Print()" function prints its arguments with their default format

	// Printing the values of "i" & "j"
	// Note -> The provided values will be printed exactly with no space between them
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)
	fmt.Print("\n")

	//To print argument in new line use "\n"
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// It is also Possible to use only one "Print()" for Printing multiple variables
	fmt.Print(i, "\n", j, "\n")

	// If we want to add a space between string arguments, we need to use " "
	fmt.Print(i, " ", j, "\n")

	// "Print()" inserts a sapce in between if neither one one of them is a string
	var a, b = 10, 20

	fmt.Print(a, b)
	fmt.Print("\n")

	// "Println()" -> The "Println()" function is similar to "Print()" with the difference that a whitespace is added between the arguments, and a newline is added at the end
	fmt.Println(i, j)

	// "Printf()" -> The "Print()" function first formats its arguments based on the given formatting verb and then prints them
	// -> "%v" -> is ised to print the 'value' of the arguments
	// -> "%T" -> is used to print the 'type' of the arguments

	var p string = "Hello"
	var q int = 15

	fmt.Printf("p has value: %v & type: %T\n", p, p)
	fmt.Printf("q has value: %v & type: %T\n", q, q)

}
