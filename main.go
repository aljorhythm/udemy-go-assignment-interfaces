package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	width  float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return t.height * t.width * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	printArea(triangle{2, 3})
	printArea(square{2})
}
