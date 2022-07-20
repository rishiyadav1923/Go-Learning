/*
	Race -> A Race Condition occurs in Go when two or more goroutines try to access the same
	resource. It may happen when a variable attempts to read & write the resource without any
	regaord to other routines
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count int

func increment(s string) {

	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count:", count)
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go increment("Foo:")
	go increment("Bar:")

	wait.Wait()

	fmt.Println("Last count Value", count)
}
