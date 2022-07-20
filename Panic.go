/*
	Go Panic is a mechanism by which we handle error situation.
	Panic can be used to abort a function execution. When a function
	calls panic, its execution stops and the control flows to the
	associated deferred function
	The caller of this function also gets terminated and callers
	deferred function gets executed (if present any). This process
	continues till the program terminates. Now the error condition
	is reported
*/

package main

import "os"

func main() {
	panic("Error Situation")
	_, err := os.Open("\temp\file")
	if err != nil {
		panic(err)
	}

}
