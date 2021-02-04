// Structs

package main

import ( 
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Returns a new Square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	square := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return square, nil
}

// Move the square
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {

	square, err := NewSquare(2, 5, 20)

	if err != nil {
		log.Fatalf("ERROR: cannot create a square")
	}

	square.Move(2, 3)

	fmt.Println("------------")
	fmt.Printf("%+v\n", square)
	fmt.Println("----------")
	fmt.Println("Area is", square.Area())
	fmt.Println("------------")

}