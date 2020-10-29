package main

import (
	"fmt"
)

func main() {
	age := 10

	if age >= 10 {
		fmt.Println("bisa maen game")
	} else {
		fmt.Println("beloman bisa")
	}

	nilai := 80
	var grade string

	if nilai <= 50 {
		grade = "E"
	} else if nilai <= 60 {
		grade = "D"
	} else if nilai <= 70 {
		grade = "C"
	} else {
		grade = "Lulus"
	}

	fmt.Println(grade)
}
