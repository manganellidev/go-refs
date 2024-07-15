package foundation

import "fmt"

func Arrys() {
	fmt.Println("\n--> Arrays")

	var EvenNum[5] int
	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8
	fmt.Println(EvenNum[2])

	EvenNumShort := [5]int{0,2,4,6,8}
	fmt.Println(EvenNumShort[2])

	// for _, value := range EvenNum {
	// 	fmt.Println(value)
	// }

	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	numSlice := []int{5,4,3,2,1}
	sliced := numSlice[3:5]
	fmt.Println(sliced)
	sliced2 := numSlice[:5]
	fmt.Println(sliced2)
	sliced3 := numSlice[0:]
	fmt.Println(sliced3)

	slice4 := make([]int, 5, 10)
	copy(slice4, numSlice)
	fmt.Println(slice4)

	slice5 := append(numSlice, 3, 0, -1)
	fmt.Println(slice5)
}