/*
	Syntax ->
	func FunctionName() {
		Code to be executed
	}
*/

package main

import (
	"fmt"
)

// Function to be called
func myMessage() {
	fmt.Println("I just got executed !")
}

func main() {
	myMessage() // call the function

	// You can also make multiple function call for same function
	myMessage()
	myMessage()
	myMessage()
}

/*	Naming rules for Go Functions
	-> A Function name must start with a letter
	-> A Function name can only contain alpha-numeric characters and underscores ( A-Z, a-z, 0-9, & _)
	-> Function names are case-sensitive
	-> A Function name cannot contain spaces
	-> If the function name consists of multiple words, techinques introduced for multi-word variable naming can be used
*/
