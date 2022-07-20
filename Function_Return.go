/*
	Syntax ->
	func FunctionName(parameter-1 type, parameter-2 type) type {
		code to be executed
		return output
	}
*/

package main

import (
	"fmt"
)

func myfunction(x int, y int) (result int) {
	result = x + y
	return
} // Storing the Returned Value in a Variable

func Myfunction(x int, y int) (result int) {

	result = x + y
	return
} // Named Return Values -> returning a naked value (Means No value is Specified)

func myFunction(x int, y int) int {

	return x + y
} // Function Returns

func main() {
	fmt.Println(myFunction(1, 2))
	fmt.Println(Myfunction(1, 2))
	total := myfunction(1, 2)
	fmt.Println(total)
}
