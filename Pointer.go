// Syntax -> var var_name *var-type

package main

import (
	"fmt"
)

func main() {

	x := 10

	changeX(&x)
	fmt.Println(x)

}

func changeX(x *int) {

	*x = 0
}
