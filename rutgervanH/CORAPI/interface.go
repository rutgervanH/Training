package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}


func (s square) area() float64{
	return s.side * s.side
}


func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}


func main() {
	a := circle{5}
	b := square{3}
	info(a)
	info(b)

fmt.Println(totalArea(a,b))

}
