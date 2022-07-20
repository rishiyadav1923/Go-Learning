// Go Reflection is the ability of a program to examine its own structure, particularly through the types; its a form of
// meta-programming

package main

import (
	"fmt"
	"reflect"
)

func main() {

	age := 27.5

	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))
}
