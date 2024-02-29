package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

type Point interface {
	X() float64
	Y() float64
}
 
func (p *point) X() float64 {
	return p.x
}

func (p *point) Y() float64 {
	return p.y
}

func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.X() - p2.X()), 2) + math.Pow((p1.Y() - p2.Y()), 2))
}

func main() {
	p1 := NewPoint(-4, 4)
	p2 := NewPoint(2, 4)
	d := Distance(p1, p2)
	fmt.Println(d)
}