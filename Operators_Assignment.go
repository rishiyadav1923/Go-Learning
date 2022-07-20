package main

import (
	"fmt"
)

func main() {

	var a int = 5
	fmt.Println(a)

	var b int = 5
	b += 3
	fmt.Println(b)

	var c int = 5
	c -= 3
	fmt.Println(c)

	var d int = 5
	d *= 3
	fmt.Println(d)

	var e int = 5
	e /= 3
	fmt.Println(e)

	var f int = 5
	f %= 3
	fmt.Println(f)

	var g int = 5
	g &= 3
	fmt.Println(g)

	var h int = 5
	h |= 3
	fmt.Println(h)

	var i int = 5
	i ^= 3
	fmt.Println(i)

	var j int = 5
	j >>= 3
	fmt.Println(j)

	var k int = 5
	k <<= 3
	fmt.Println(k)

}
