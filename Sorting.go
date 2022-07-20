/*
	Go has sort package which can be used for sorting the built-in as well as user defined data types.
	The sort package has different methods to sort different data types like Ints(), Float64s(), Strings() etc.
	We can check if the values are sort or not by using the AreSorted() methods like Float64sAreSorted(), IntsAreSorted() etc.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	intValue := []int{10, 20, 5, 8}
	sort.Ints(intValue)
	fmt.Println("Ints:   ", intValue)

	floatValue := []float64{10.5, 20.5, 5.5, 8.5}
	sort.Float64s(floatValue)
	fmt.Println("floatValue:   ", floatValue)

	stringValue := []string{"Raj", "Mohan", "Roy"}
	sort.Strings(stringValue)
	fmt.Println("Strings:", stringValue)

	str := sort.Float64sAreSorted(floatValue)
	fmt.Println("Sorted: ", str)
}
