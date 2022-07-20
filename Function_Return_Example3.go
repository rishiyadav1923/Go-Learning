// Multiple return values
// Here, we want to omit the second returned value (txt1 -> which is stored in variable b)

package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {

	a, _ := myFunction(5, "Hello")
	fmt.Println(a)
}
