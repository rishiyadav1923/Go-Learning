/*
	Syntax ->
	type struct_name struct {
		member1 datatype;
		member2 datatype;
		member3 datatype;
*/

package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	job string
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

	// Access and print pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age; ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age; ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

}
