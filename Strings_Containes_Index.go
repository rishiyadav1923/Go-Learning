package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hi...there"

	fmt.Println(strings.Contains(str, "th")) //Containes()

	str1 := "Hi...there"

	fmt.Println(strings.Index(str1, "th")) //Index()
}
