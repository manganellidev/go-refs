package foundation

import "fmt"

func Recursion() {
	fmt.Println("\n--> Recursion")
	num := 5
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		fmt.Println("here")
		return 1
	}

	return num * factorial(num-1)
}