package main

import (
	"log"
	"time"
)

func haveFun(s string) {

	log.Printf("\tA: Let's have fun: %v", s)
}

func doPolling() {
	for _ = range time.Tick(2 * time.Second) {
		haveFun("\t B: Okay!")
	}
}

func main() {

	go doPolling()

	select {} // The select statement lets a Goroutine wait on multiple communication operations
}
