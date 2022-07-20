/*
	Concurrency -> Large parameters are divided into smaller sub-programs. Programs which run their smaller
	components at the same time is known as concurrency
*/
/*
	Goroutine -> The parts of an application that run concurrently are called goroutines. Goroutines and channels
	are used for structuring concurrent programs.
	A Process is an machine which runs in its own address space in memory. A process has threads which are simultaneously
	executing entities. Threads share the same address space of the process

	Goroutines are Lightweight, much lighter than thread. Goroutine run in the same address space, so acccess
	to shared memory must be synchronized; This can be done by sync package, but it is recommended to use channels
	to synchronize goroutines.

	A Goroutine is implementd as a function or ethod. It is called (Invoke) with the "Go" keyword. When the goroutine
	finishes, nothing is returned to the caller function.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go fun1()
	go fun2()
	wg.Wait()
}

func fun1() {

	for i := 0; i < 10; i++ {
		fmt.Println("fun1, ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

func fun2() {

	for i := 0; i < 10; i++ {
		fmt.Println("fun2, ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
