// Multiple Return Values
// Here, we store the two return values into two variables (a & b)

package main

import (
	"fmt"
)

func function(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + "World!"
	return
}

func main() {

	a, b := function(5, "Hello ")
	fmt.Println(a, b)

}
