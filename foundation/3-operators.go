package foundation

import "fmt"

func Operators() {
	fmt.Println("\n--> Operators")
	// Arithmetic
	// + addition
	// - substraction
	// * multiplication
	// / division
	// % modulus
	x,y := 6,3
	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y)
	fmt.Println("x mod y =", x%y)

	// Relational
	// > greater than
	// < lesser than
	// >= greater than or equal
	// == equivalence
	// != not equal
	fmt.Println(func() string {
		if x > y {
			return "x is greater than y"
		} else {
			return "y is greater than or equal to x"
		}
	}())

	// Logical
	// && and
	// || or
	// ! negation
	fmt.Println(func() bool {
		return x == y && x > 0
	}())
}