/*
 A Closure is a Function which refers reference variable from outside its body. The function
 may access and assign to the referenced variables.
*/

package main

import (
	"fmt"
)

func main() {

	number := 10

	squareNum := func() int {
		number *= number
		return number
	}

	fmt.Println(squareNum())
	fmt.Println(squareNum())
}
