package main

type Point struct {
	x float64
	y float64
}

func Add(a, b Point) Point {

	return Point{a.x + b.x, a.y + b.y}

}

func NewPoint(a, b float64) Point {

	return Point{a, b}
}
