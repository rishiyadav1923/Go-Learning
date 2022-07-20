// Recover is used to regain control of a program from panic or error-condition.

package main

import (
	"fmt"
)

func main() {

	fmt.Println(SaveDivide(10, 0))
	fmt.Println(SaveDivide(10, 10))
}

func SaveDivide(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	quotient := num1 / num2
	return quotient
}
