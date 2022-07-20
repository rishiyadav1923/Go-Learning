package main

import (
	"fmt"
	"time"
)

func main() {

	present := time.Now()

	fmt.Println("Today :", present.Format("Mon ,Jan 2, 2006 at 3:04pm"))
	someTime := time.Date(2017, time.March, 30, 11, 30, 55, 123456, time.Local)

	// Compare time with time Equal()
	sameTime := someTime.Equal(present)

	fmt.Println("SomeTime equals to now ? :", sameTime)

	// Calculate the time different between today
	// And long time ago

	diff := present.Sub(someTime)

	// Convert diff to days

	days := int(diff.Hours() / 24)
	fmt.Printf("30th March 2017 was %d days ago \n", days)
}
