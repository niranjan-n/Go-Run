package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.breadth
}
func (c *Circle) Area() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rectangle{20, 10}
	c := Circle{2}
	val1, val2 := r.Area(), c.Area()
	fmt.Println("Area of Rec:", val1, "Area of Circle :", val2)
}
