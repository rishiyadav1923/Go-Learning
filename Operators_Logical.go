package main

import (
	"fmt"
)

func main() {

	var a int = 5

	fmt.Println(a < 5 && a < 10)
	fmt.Println(a < 5 || a < 10)
	fmt.Println(!(a < 5 && a < 10))
}
