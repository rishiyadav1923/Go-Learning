package main

import (
	"fmt"
)

func main() {

	/*
		Syntax ->
		if (condition-1) {
			code to be executed if condition-1 is true
		}
		if (condition-2) {
			code to be executed if both condition-1 & condition-2 are true
		}
	*/

	num := 20

	if num >= 10 {
		fmt.Println("Num is more than 10")
	}
	if num > 15 {
		fmt.Println("Num is also more than 15")
	} else {
		fmt.Println("Num is less than 10")
	}
}
