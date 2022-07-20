/*
	Mutex -> Mutual Exclusion locks, or mutexes can be used to synchronize access to state
	& safely access data across many goroutines. It acts as a guard to the entrance of the
	critical section of code so that only one thread can enter the critical section at a time

	We set a Lock around Particular lines of code with it. While one Goroutine holds the lock,
	all other Goroutines are prevented from executing any lines of code protected by the same
	mutex, and are forced to wait until the lock is yielded before they can proceed
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
var mutex sync.Mutex

func increment(s string) {

	for i := 0; i < 10; i++ {
		mutex.Lock()
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count:", count)
		mutex.Unlock()
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
