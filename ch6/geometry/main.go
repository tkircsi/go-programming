package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func main() {
	p := Point{1, 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	p.ScaleBy(3.0)
	fmt.Println(p)

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
