package main

import (
	"fmt"
)

func main() {
	//var myMap map[string]int{}
	myMap := map[string]int{}

	myMap["ruby"] = 10000
	myMap["ps5"] = 90000000
	fmt.Println(myMap)

	var nama string = "utha"

	fmt.Println(nama)

	//cara manggil suatu nilai map
	//klo di array kan kita bisa 	fmt.Println(nama[0])
	//klo di map kita bisa panggil indexnya / key nya
	fmt.Println(myMap["ps5"])

	//Map bisa juga dibikin kaya aray bikin variabel lg bisa di isi valuenya
	myGame := map[int]string{1: "PS3", 2: "PS5", 3: "PS1"}
	fmt.Println(myGame)

	//inget klo vertikal pake koma di trakhirnya
	myGame2 := map[int]string{
		1: "PS3",
		2: "PS5",
		3: "PS1",
	}
	fmt.Println(myGame2)

	//for each di map
	for key, result := range myGame2 {
		fmt.Println("key nya =", key, "valuenya =", result)
	}

	fmt.Println("======================================")
	//delete di map dengan keynya
	delete(myGame2, 1)
	//for each di map
	for key, result := range myGame2 {
		fmt.Println("key nya =", key, "valuenya =", result)
	}
}
