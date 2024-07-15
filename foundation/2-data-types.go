package foundation

import "fmt"

func DataTypes() {
	fmt.Println("\n--> Data Types")
	// Numeric
	var age int = 10
	var price float32 = 3.367
	fmt.Println("age:", age)
	fmt.Println("price:", price)

	// Boolean
	var isBool bool = true
	fmt.Println("isBool:", isBool)

	// String
	var name string = "Wagner Manganelli"
	fmt.Println("name:", name)

	// Derived
	// - Pointer
	// - Array
	// - Structure
	// - Map
	// - Interface
}
