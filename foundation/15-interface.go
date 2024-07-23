package foundation

import (
	"fmt"
	"math"
)

func Interfaces() {
	fmt.Println("\n--> Interfaces")

	rect := Rectangle2{50,60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is", getArea(&rect))
	fmt.Println("Area of circle is", getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle2 struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle2) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius,2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}