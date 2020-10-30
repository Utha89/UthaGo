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
}
