package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, heigth float64
}

type circle struct {
	rain float64
}

func (s square) area() float64 {
	return s.width * s.heigth
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.heigth
}

func (c circle) area() float64 {
	return math.Pi * c.rain * c.rain
}

func (c circle) perim() float64 {
	return (2 * math.Pi * c.rain)
}

func meansure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	square := square{width: 3, heigth: 4}
	circle := circle{rain: 5}

	meansure(square)
	meansure(circle)
}
