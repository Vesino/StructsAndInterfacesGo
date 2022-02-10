package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// function
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Method <3
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c)) // using a function

	// This is much easier to read. We no longer need the & operator
	// (Go automatically knows to pass a pointer to the circle for this method)
	fmt.Println(c.area()) // using a method
}
