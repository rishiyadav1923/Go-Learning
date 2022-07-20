package main

import (
	"fmt"
	"strings"
)

func main() {

	var arr = []string{"a", "b", "c", "d"}

	fmt.Println(strings.Join(arr, "*")) //Join()

	var str = "New"

	fmt.Println(strings.Repeat(str, 4)) //Repeat()
}
