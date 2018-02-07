package main

import ("fmt")

func varrie (n ...int) (int,int) {
	var l int
	var ni int
	for i, v := range n {
		if v > l {
			l = v
			ni = i
		}
	}
	return l, ni
}

func main () {

	fmt.Println(varrie(2,4,6,8,9,0))
}
