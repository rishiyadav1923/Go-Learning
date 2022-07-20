// Slices
// -> Using the "[]datatype{values} format"
// -> Create a Slice from an array
// -> Using the "make()" function

// len() function -> returns the length of the slice (the number of elements in the slice)
// cap() function -> returns the capacity of the slice (the number of elements the slice can grow or shrink to)

// SYNTAX FOR "[]datatypes{value}"
// Syntax :- slice_name := []datatype{values}
// example - myslice := []int
// example - myslice := []int{1,2,3}

// SYNTAX FOR "Creating a Slice from an Array"
// Syntax :- var myarray = [length]datatype{values}	// An array
// Syntax :- myslice := myarray[start:end]	// A slice made from the array

// SYNTAX FOR "make() function"
// slice_name := make([]type, length, capacity)

package main

import (
	"fmt"
)

func main() {

	// -> "[]datatype{values}"
	myslice1 := []int{}

	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"GO", "Slices", "Are", "Powerfull"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	fmt.Println()

	// -> Creating Slice from an array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("Length = %d\n", len(myslice))
	fmt.Printf("Capacity = %d\n", cap(myslice))

	// -> Creating Slice with "make()" function
	// -> Without Omitted capacity
	mySlice1 := make([]int, 5, 10)

	fmt.Printf("mySlice = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %d\n", cap(mySlice1))

	// -> With Omitted capacity
	mySlice2 := make([]int, 5)
	fmt.Printf("mySlice2 = %v\n", mySlice2)
	fmt.Printf("Length = %v\n", len(mySlice2))
	fmt.Printf("Capacity = %v\n", cap(mySlice2))
}
