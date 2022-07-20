package main

import (
	"fmt"
)

func main() {

	/*
		if(condition) {
			 code to be executed if condition is true
		} else {
			 code to be executed if condition is false
		}
	*/

	time := 20

	// This will execute if statement

	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening")
	}

	temperature := 14

	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}
}
