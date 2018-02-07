package main

import (
	"fmt"
)



func main() {


	dodo := func (z int) (int, bool) {

	return z / 2, z%2 == 0

	}

	fmt.Println("hallo world")
	fmt.Println(dodo(6))

	val, tr := dodo(7)
	fmt.Println(val, tr)

}
