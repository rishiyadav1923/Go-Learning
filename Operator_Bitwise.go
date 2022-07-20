package main

import (
	"fmt"
)

func main() {

	var a = 9
	var b = 8

	fmt.Printf("a = %b\n", a)
	fmt.Printf("b = %b\n", b)
	fmt.Printf("a & b is %b\n", a&b)
	fmt.Printf("a OR b is %b\n", a|b)
	fmt.Printf("a XOR b is %b\n", a^b)
	fmt.Printf("a Zero fill left shift b is %b\n", a<<b)
	fmt.Printf("a Signed right Shift b is %b\n", a>>b)
}
