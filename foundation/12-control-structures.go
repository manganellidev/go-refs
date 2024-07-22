package foundation

import "fmt"

func ControlStructures() {
	fmt.Println("\n--> Control Structures: Defer, Recover, and Panic")
	
	defer firstRun()
	secondRun()

	fmt.Println(div(3,0))
	fmt.Println(div(6,2))

	demPanic()
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic!!!")
}

func firstRun() {
	fmt.Println("I executed first!")
}

func secondRun() {
	fmt.Println("I executed second!")
}