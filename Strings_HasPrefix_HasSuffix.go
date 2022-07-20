package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "INDIA"

	fmt.Println(strings.HasPrefix(str, "IN")) //HasPrefix()
	fmt.Println(strings.HasSuffix(str, "IA")) //HasSuffix()
}
