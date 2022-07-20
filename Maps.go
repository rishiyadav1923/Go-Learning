/*
	Syntax ->
	var a = map[keyType]ValueType{key1:value1, key2:value2,...}
							OR
	b := map[keyType]ValueType{key1:value1, key2:value2,...}
*/
// Maps are used to store data values in key:value pairs
// Each elemet in a map is a key:value pair
// A amp is an unordered and changeable collection that does not allow duplicates
// The Length of a map is the number of its elements. You can find it using the len() Function
// The dafault value of a map is nil
// Maps hold references to an underlying hash table
// Go has multiple ways for creating maps

package main

import (
	"fmt"
)

func main() {

	var a = map[string]string{"brand": "Ford", "Model": "Mustang", "Year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

// NOTE -> The order of the map elements defined in the code is differernt from the way
// that they are stored. The data are stored in a away to have efficient data retrieval from
// the map.
