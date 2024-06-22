package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x int
	y int
}

func New(x int, y int) Coordinates {
	return Coordinates{
		x: x,
		y: y,
	}
}

func (c Coordinates) findDistance() int {
	return int(math.Abs(float64(c.x - c.y)))
}

func main() {
	c := New(2, 7)
	fmt.Println(c.findDistance())
}
