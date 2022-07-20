// Integer Formatting Verbs
// %b -> Base 2
// %d -> Base 10
// %+d -> Base 10 and always show sign
// %o -> Base 8
// %O -> Base 8, with leading 0o
// %x -> Base 16, lowercase
// %X -> Base 16, lowercase
// %#x -> Base 16, with leading 0x
// %4d -> Pad with spaces (width 4, right justified)
// %-4d -> Pad with spaces (width 4, left justified)
// %04d -> Pad with zeroes (width 4)

package main

import (
	"fmt"
)

func main() {
	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
}
