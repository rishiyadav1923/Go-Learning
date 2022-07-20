// Multiple Return Values
// Here, myFunction() returns one integer (result) and one string (txt1)

package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + "World!"
	return
}

func main() {
	fmt.Println(myFunction(5, "Hello"))
}
