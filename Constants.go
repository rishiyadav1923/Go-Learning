// If a variable should have a fixed value that cannot be changed, you can use the "const" keyword
// Syntax -> const CONSTNAME type = value
// The value of a constant must be assigned when you declare it

package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	fmt.Println(PI)
}

// Rule 1 -> Constant names follow the same naming rules as variable
// Rule 2 -> Constant names are usually written in uppercase letters (for easy identification & differentiation from variables)
// Rule 3 -> Constant can be declared both inside and outside of a function
