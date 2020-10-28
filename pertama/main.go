package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hello World")
	//cara panggil function yg masih sama2 1 pacake yaitu package main
	kalimat := TestAja()

	fmt.Println(kalimat)

	//cara panggil enjumlahan 2 buah angka
	angka := calculation.Add(8, 9)

	fmt.Println(angka)
}
