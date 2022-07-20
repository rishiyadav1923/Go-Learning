/*
	Atomic Variable are used to manage state, though sync/atomic package and avoid race
	condition. Atomic counter can be accessed by multiple go routines.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wait sync.WaitGroup
var count int64

func increment(s string) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&count, 1)
		fmt.Println(s, i, "Count ->", count)
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
