// Arrays -> Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value
// Syntax -> Using "var" keyword :-
// var array_name = [length]datatype{value} // here length is defined
// var array_name = [...]datatype{values} // here length is inferred
// Syntax -> Using ":=" sign :-
// array_name := [length]datatype{values} // here length is defined
// array_name := [...]datatype{values} // here length is inferred

package main

import (
	"fmt"
)

func main() {

	// Declaring 2 arrays (arr1 & arr2) with defined lengths:
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println("Declaring 2 arrays with defined length")
	fmt.Println(arr1)
	fmt.Println(arr2)

	// Declaring 2 arrays (arr3 & arr4) with inferred length:
	var arr3 = [...]int{1, 2, 3}
	var arr4 = [...]int{4, 5, 6, 7, 8}

	fmt.Println("Declaring 2 arrays with inferred length")
	fmt.Println(arr3)
	fmt.Println(arr4)

	// Declaring an arrays of "String"
	var cars = [4]string{"Volvo", "Tesla", "Mercedez", "BMW"}

	fmt.Printf("Declared array of strings are: \n")
	fmt.Println(cars)

	// Accessing Elements of an Arrays
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0]) // The values present in "[]" are index of the number present in the arrays
	fmt.Println(prices[2])

	// Changing Elements of an Array
	cost := [3]int{10, 20, 30}

	cost[2] = 50
	fmt.Println(cost)

	// Array Initialization
	arr5 := [5]int{}              // not initialized
	arr6 := [5]int{1, 2}          // partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // Fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	// Initializing only Specific Elements
	arr8 := [5]int{1: 10, 2: 40}

	// -> "1:10" means assign 10 to array index 1 (second element)
	// -> "2:40" means assign 40 to array index 2 (third element)
	fmt.Println(arr8)

	// Finding length of an Array
	// the "len()" function is used to find the length of an array:

	arr9 := [4]string{"Volvo", "BMW", "Ford", "Mercedez"}
	arr10 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr9))
	fmt.Println(len(arr10))
}
