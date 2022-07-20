package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hi...There"

	fmt.Println(strings.Count(str, "e")) //Count()

	str1 := "Hi...There"

	fmt.Println(strings.Replace(str1, "e", "Z", 2)) //Replace
}
