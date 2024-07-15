package foundation

import "fmt"

func Loops() {
	for i:=1; i<=10; i++ {
		fmt.Print(i," ")
	}

	println()

	i := 1
	for i <= 10 {
		fmt.Print(i," ")
		i++
	}

	println()

	for i:=1; i<=4; i++ {
		for j:=1; j<i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}