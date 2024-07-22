package foundation

import "fmt"

func Structures() {
	fmt.Println("\n--> Structures")
	react1 := Rectangle{height: 10, width: 5}
	react2 := Rectangle{10, 5}
	fmt.Println(react1)
	fmt.Println(react2)
	fmt.Println(react2.height)
	fmt.Println(react2.width)
	fmt.Println("Area of the rectangle is", react2.area())
}

type Rectangle struct {
	height float64
	width float64
}

func (react *Rectangle) area() float64 {
	return react.height * react.width
}