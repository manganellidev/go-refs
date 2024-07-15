package foundation

import "fmt"

func VariablesConstants() {
	fmt.Println("\n--> Variables Constants")
	// A constant is a variable with a value that can't be changed
	const pi float64 = 3.14159265359
	fmt.Println("pi:", pi)

	// You can decalre multiple variables like this
	var (
		varA = 2
		varB = 3
	)

	// for short
	x,y := 10,15
	fmt.Println(x,y)

	fmt.Println(varA, varB)

	// Strings are a series of characters between " or  `
	var name string = "Wagner Manganelli"
	fmt.Println(len(name))
}