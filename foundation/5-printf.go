package foundation

import "fmt"

func Printf() {
	fmt.Println("\n--> Printf")

	name,lastName := "Wagner","Manganelli"
	fmt.Printf("my name is %s %s \n", name, lastName)
	fmt.Printf("%T \n",name)

	price := 3.12
	fmt.Printf("the price is %.2f \n", price)
	fmt.Printf("%T \n",price)

	win := true
	fmt.Printf("%t \n",win)

	age := 31
	fmt.Printf("%d \n",age)

	fmt.Printf("%b \n",25)
	fmt.Printf("%c \n",34)
	fmt.Printf("%x \n",15)
	fmt.Printf("%e \n",price)
}