// Interfaces

package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Return the sum of all areas

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// Shape is an interface, a collection of methods which must be implemented i
type Shape interface {
	Area() float64
}


func main() {
	s := &Square{20}
	fmt.Println("------------")
	fmt.Println("Square area: ", s.Area())

	c := &Circle{10}
	fmt.Println("------------")
	fmt.Println("Circle area: ", c.Area())

	shapes := []Shape{s, c}

	total := sumAreas(shapes)
	fmt.Println("------------")
	fmt.Println("Total area: ", total)

	fmt.Println("------------")
}