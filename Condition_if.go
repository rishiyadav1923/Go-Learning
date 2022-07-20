package main

import (
	"fmt"
)

func main() {

	/*
		Syntax -> if (condition) {
			code to be executed if condition is true
		}
	*/

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	x := 20
	y := 18

	if x > y {
		fmt.Println("x is greater than y")
	}
}
