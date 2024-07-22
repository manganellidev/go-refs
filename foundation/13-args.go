package foundation

import "fmt"

func Args() {
	fmt.Println("\n--> Args")

	fmt.Println(sumMany(10,20,30,40,50))
}

func sumMany(args ... int) int {
	sum := 0

	for _, value := range args {
		sum += value
	}

	return sum
}