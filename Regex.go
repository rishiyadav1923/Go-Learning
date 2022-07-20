// Go Regex is Used for searching String
package main

import (
	"fmt"
	"regexp"
)

func main() {

	re := regexp.MustCompile(".com")

	// To find string having the text left most match . If no match is found , empty string is returned
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.com"))
	fmt.Println(re.FindString("fb.com"))

	//To find index value fo the left most text . If no match found then empty brackets are returned
	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("abc.com"))
	fmt.Println(re.FindStringIndex("fb.com"))

	//To return the slice of string having the text of the leftmost match and the matches , we use "FindStringSubmatch()"
	re1 := regexp.MustCompile("f[a-z]+ing")

	fmt.Println(re1.FindStringIndex("google.com"))
	fmt.Println(re1.FindStringSubmatch("flying"))
	fmt.Println(re1.FindStringSubmatch("abcfloatingxyz"))

}
