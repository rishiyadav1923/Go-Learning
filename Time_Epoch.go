/*
	Epoch -> I Go, We can produce the time in seconds, milliseconds, nanoseconds.
	The time package has the required methods like secs, nanos etc which help us.

	The reference time is the unix Epoch. We may also convert nanoseconds or
	milliseconds into time format.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	current_time := time.Now()
	secs := current_time.Unix()
	nanos := current_time.UnixNano()

	fmt.Println(current_time)

	millis := nanos / 1000000
	p(secs)
	p(millis)
	p(nanos)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
