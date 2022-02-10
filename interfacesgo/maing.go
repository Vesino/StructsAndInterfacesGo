package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// Method <3
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	b, h float64
}

func (r *Rectangle) area() float64 {
	return r.b * r.h
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{5, 5}
	fmt.Println(totalArea(&c, &r))
}
