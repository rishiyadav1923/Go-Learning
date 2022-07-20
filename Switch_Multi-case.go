package main

import (
	"fmt"
)

func main() {

	/*
		Syntax ->
		switch (expression) {
		case x,y:
			code block if expression is evaluted to x or y
		case v,w:
			code block if expression is evaluated to v or w
		case z:
			...
		dafault:
			code block if expression is not found in any cases
		}
	*/

	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd Weekday")
	case 2, 4:
		fmt.Println("Even Weekday")
	case 6, 7:
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day of day number")
	}
}
