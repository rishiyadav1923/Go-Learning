// Syntax -> var a map[KeyType]ValueType
// The "make()" function is the right way to create an empty map.
// If you make an empty map in a different way and write to it, it will causues a runtime Panic.

package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
