// Syntax -> const identifier [type] = value
// Example -> const PI = 3.14159
// Explicit Type -> const b string = "abc"
// Implicit Type -> const b = "abc"

package main

import (
	"fmt"
)

func main() {

	const HEIGHT int = 100
	const WIDTH int = 200

	var area int

	area = HEIGHT * WIDTH

	fmt.Printf("Value of area: %d", area)
}
