package main

import (
	"fmt"
)

func main() {

	/*
		Syntax ->

		if (condition-1) {
			code to be executed if condition-1 is true
		} else if (condition-2) {
			code to be executed if condition-1 is false and condition-2 is true
		} else {
			code to be executed if condition-1 & condition-2 are both false
		}
	*/

	time := 22

	if time < 10 {
		fmt.Println("Good Morning")
	} else if time < 20 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

	a := 14
	b := 14

	if a < b {
		fmt.Println("a is less than b")
	} else if a < b {
		fmt.Println("a is more than b")
	} else {
		fmt.Println("a & b are equal")
	}

	x := 30

	if x >= 10 {
		fmt.Println("x is larger than or equal to 10")
	} else if x > 20 {
		fmt.Println("x is larger than 20")
	} else {
		fmt.Println("x is less than 10")
	}
}
