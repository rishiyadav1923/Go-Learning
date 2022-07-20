package main

import (
	"fmt"
)

//<--Varible Declaration with Initial Value -->//
// int -> stores integers (whole numbers), such as 123 or -123
// float32 -> stores floating point numbers, with decimals, such as 19.99 or -19.99
// string -> stores text, such as "Hello World". String values are surrounded by double quotes
// bool -> stores values with two states: true or false

func main() {
	//Syntax -> var variablename type = value
	//Syntax -> variablename := value

	var student1 string = "Rishi" // type is string
	var student2 = "Sidhu"        // type is inferred

	x := 2 // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
