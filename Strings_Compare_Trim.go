package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a")) //	Compare()
	fmt.Println(strings.Compare("b", "a"))

	fmt.Println(strings.TrimSpace("\t\n I Love My Country \n\t\r\n")) // Trim()
}
