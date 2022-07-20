/*
	WorkerPool -> WorkerPool is a design in which a fixed number of M Workers (Go Routines) works on
	N Taks in a Work queue (Go channel). The Work resides in a queue until a worker finish its
	current task and pull a new one
*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "Processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	job := make(chan int, 10)
	result := make(chan int, 10)

	for w := 1; w <= 2; w++ {
		go worker(w, job, result)
	}

	for j := 1; j <= 9; j++ {
		job <- j
	}

	close(job)

	for a := 1; a <= 9; a++ {
		<-result
	}
}
