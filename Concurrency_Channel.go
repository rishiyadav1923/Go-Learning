/*
	Channel -> The Channel acts as a pipe by which we send typed values from one Goroutine to Another
	It guarantees synchronization since only one Goroutine has access to a data item at any given time.
	The Ownership of the data is passed between different Goroutine. Hence, By design it avoids the pitfalls
	of shared memory & prevent face condition
*/

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true

}

func main() {

	done := make(chan bool, 1)
	go worker(done)
	<-done
}
