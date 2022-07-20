/*
	Go Tickers are used when we want to some work at regular interval of time. Tickers can be stopped
	like timers using the Stop() method.

	The NewTicker() method returns a new Ticker having a channel which send the time according to the
	duration argument. The duration must be larger than zero, if not, the ticker will panic.

	The Tick() is a wrapper for NewTicker which provides access to the ticking channel. The Tick()
	method is useful for clients who dont want to shutdown the Ticker.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	tickerValue := time.NewTicker(time.Millisecond * 100)

	go func() {

		for t := range tickerValue.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 500)
	tickerValue.Stop()
	fmt.Println("Ticker Stopped")

}
