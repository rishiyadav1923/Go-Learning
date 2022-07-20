// Multiple return values
// Here, we want to omit the first returned value (result -> which is stored in variable a)
// If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value

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

	_, b := myFunction(5, "Hello ")
	fmt.Println(b)
}
