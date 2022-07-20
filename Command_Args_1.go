package main

import (
	"fmt"
	"os"
)

func main() {

	argumentWithPath := os.Args  // Returns all arguments including path
	argumentSlice := os.Args[1:] // Returns all elements after path
	argumentAt2 := os.Args[2]    // Returns specified argument only
	fmt.Println(argumentWithPath)
	fmt.Println(argumentSlice)
	fmt.Println(argumentAt2)
}
