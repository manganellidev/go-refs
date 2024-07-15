package foundation

import "fmt"

func Pointers() {
	fmt.Println("\n--> Pointers")
	x := 5

	fmt.Println(x)
	fmt.Println(&x)

	changeValue(x)
	fmt.Println("without pointer:", x)
	changeValueWithRef(&x)
	fmt.Println("with pointer:", x)
}

func changeValue(x int) {
	x = 7
}

func changeValueWithRef(x *int) {
	*x = 7
}