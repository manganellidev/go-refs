package foundation

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Files() {
	fmt.Println("\n--> Files")

	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, my name is Wagner and this text was created using Go!")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)
	fmt.Println(s1)
}