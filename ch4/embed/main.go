package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 14}, 20}

	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Point:  Point{2, 5},
			Radius: 4,
		},
		Spokes: 40,
	}

	fmt.Printf("%#v\n", w)

	// Struct Anonymous fields
	w.X = 10
	fmt.Printf("%#v\n", w)

}
