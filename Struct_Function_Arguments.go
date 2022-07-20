// Passing Struct as Function Arguments
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {

	var pers1 Person
	var pers2 Person

	// pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// pers2 specification
	pers2.name = "Rishi"
	pers2.age = 24
	pers2.job = "Devops Engineer"
	pers2.salary = 25000000

	// Print pers1 info by calling a function
	printPerson(pers1)
	// Print Pers2 info by calling a function
	printPerson(pers2)
}

func printPerson(pers Person) {

	fmt.Println("Name: ", pers.name)
	fmt.Println("Age; ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)

}
