/*
	We can also implement our own sorting schema, suppose we want
	to sort an array of string according to its length.
	In order to do that we have to implement our own Less,
	Len and Swap methods defined in the sort Interface.
	Then we have to cast our array into the implemented type.
*/

package main

import (
	"fmt"
	"sort"
)

type OrderByLengthDesc []string

func (s OrderByLengthDesc) Len() int {
	return len(s)
}
func (str OrderByLengthDesc) Swap(i, j int) {
	str[i], str[j] = str[j], str[i]
}
func (s OrderByLengthDesc) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
func main() {
	city := []string{"New York", "London", "Washington", "Delhi"}
	sort.Sort(OrderByLengthDesc(city))
	fmt.Println(city)
}
