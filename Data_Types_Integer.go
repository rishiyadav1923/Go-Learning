// Signed Integers -> Can store both positive and Negative values
// Type -> Size -> Range
// int	-> Depends on platform (32-bit or 64-bit)  -> -2147483648 to 2147483648 in 32-bit & -9223372036854775808 to 9223372036854775807
// int8	-> 8 bits/1 byte -> -128 to 127
// int16 -> 16 bits/2 byte -> -32768 to 32767
// int32 -> 32 bits/4 byte -> -2147483648 to 2147483647
// int64 -> 64 bits/8 byte -> -9223372036854775808 to 92233720368547775807

// Unsigned Integer -> Can only store non-negative values
// Type -> Size -> Range
// uint -> Depends on platform (32-bit or 64-bit) -> 0 to 4294967295 in (32-bit) & 0 to 18446744073709551615 in (64-bit)
// uint8 -> 8bits/1 byte -> 0 to 255
// uint16 -> 16bits/2 byte -> 0 to 65535
// uint32 -> 32bits/4 byte -> 0 to 4294967295
// uint64 -> 64bits/8 byte -> 0 to 18446744073709551615

package main

import (
	"fmt"
)

func main() {

	// Signed Integer
	var x int = 500
	var y int = -4500

	fmt.Println("This is Signed Integer")
	fmt.Printf("Type: %T, value: %v \n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	// Unsigned Integer
	var p uint = 500
	var q uint = 4500

	fmt.Println("This is Unsigned Integer")
	fmt.Printf("Type: %T, value: %v \n", p, p)
	fmt.Printf("Type: %T, value: %v \n", q, q)

}
