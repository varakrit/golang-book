package main

import (
	"math"
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r *c.r
}

func (c *Circle) changeRadius(r float64) {
	c.r = r
}

func main() {
	littleC := Circle{0, 0, 5}
	fmt.Println("littleC", littleC.area())
	littleC.changeRadius(10)
	fmt.Println("littleC", littleC.area())

	bigC := &Circle{0, 0, 5}
	fmt.Println("bigC", bigC.area())
	bigC.changeRadius(10)
	fmt.Println("bigC", bigC.area())
}