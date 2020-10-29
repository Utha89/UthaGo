package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println("angka", i)
	}

	//for each
	title := "utha ganteng"
	for index, letter := range title {
		fmt.Println("index", index, "letter", letter)
	}

}
