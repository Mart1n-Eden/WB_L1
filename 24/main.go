package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(o *Point) float64 {
	dx := p.x - o.x
	dy := p.y - o.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var x,y float64

	fmt.Println("Введите координаты первой точки")
	fmt.Scan(&x,&y)

	p1 := NewPoint(x,y)
	
	fmt.Println("Введите координаты второй точки")
	fmt.Scan(&x,&y)

	p2 := NewPoint(x,y)

	fmt.Println("Расстояние = ",p1.Distance(p2))
}