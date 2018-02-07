package main

import (
	"fmt"
)

func dothings(z int) (int, bool) {

	return z / 2, z%2 == 0

}

func main() {
	fmt.Println("hallo world")
	fmt.Println(dothings(4))

	val, tr := dothings(5)
	fmt.Println(val, tr)

}
