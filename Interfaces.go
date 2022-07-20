// Go has different approaches to implement the concepts of Object-Orienteation.
// Go does not have classes & inheritance.
// Go fullfill these requirements through its powerful interface.

// Syntax ->
// type Namer interface {
//			Method1(param_list)return_type
//			Method2(param_list)return_type
//			...
//}

package main

import (
	"fmt"
)

type vehicle interface {
	accelerate()
}

func foo(v vehicle) {
	fmt.Println(v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelerating?")
}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am Toyota, I accelerate Fast?")
}

func main() {

	c1 := car{"Suzuki", "Blue"}
	t1 := toyota{"Toyota", "Red", 100}

	c1.accelerate()
	t1.accelerate()

	foo(c1)
	foo(t1)
}
