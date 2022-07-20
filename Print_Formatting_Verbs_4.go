// Float Formatting Verbs
// %e -> Scientific Notation with 'e' as exponent
// %f -> Decimal point, no exponent
// %.2f -> Default width, precision 2
// %6.2f -> Width 6, precision 2
// %g -> Exponent as needed, only necessary digits
package main

import (
	"fmt"
)

func main() {

	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)
}
