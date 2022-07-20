package main

import (
	"fmt"
)

func main() {

	/*
		Syntax ->
		for (statement-1; statement-2; statement-3) {
			code to be executed for each iteration
		}
	*/

	// Case-1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Case-2
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	// Case-3 (Continue)
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// Case-4 (Break)
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// Case-5 (Nested)
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"Apple", "Orange", "Banana"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// Range Keyword

	/*
		Syntax ->
		for index, value := array|slice|map {
			code to be executed for each iteration
		}
	*/

	Citrus := [3]string{"apple", "Orange", "banana"}

	for idx, val := range Citrus {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// Tip -> To only show the value or the index, you can omit the other output using an Underscore (_)

	Doctor := [3]string{"apple", "Orange", "Banana"}
	for _, term := range Doctor {
		fmt.Printf("%v\n", term)
	}

	Juices := [3]string{"apple", "Orange", "banana"}

	for idx, _ := range Juices {
		fmt.Printf("%v\n", idx)
	}
}
