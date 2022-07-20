package main

import (
	"fmt"
	"strings"
)

func main() {

	// Example-1

	str := "I,Love,My,Country"

	var arr []string = strings.Split(str, ",") //Split()

	fmt.Println(len(arr))

	for i, v := range arr {

		fmt.Println("Index:", i, "Value:", v)
	}

	// Example-2

	fmt.Printf("%q\n", strings.Split("x,y,z", ","))
	fmt.Printf("%q\n", strings.Split("John and Jack and Johnny and Jinn", "and"))
	fmt.Printf("%q\n", strings.Split("abc", ""))
	fmt.Printf("%q\n", strings.Split("", "Hello"))

}
