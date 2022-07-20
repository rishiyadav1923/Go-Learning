// The Defer Keyword is used for cleaning purpose.
// The defer keyword postpones the execution of a
// function or statement until the end of the calling function

package main

import (
	"fmt"
)

func main() {

	defer print1("Hi...")
	print2("There")

}

func print1(s string) {
	fmt.Println(s)
}

func print2(s string) {
	fmt.Println(s)
}
