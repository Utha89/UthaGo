package main

import (
	"fmt"
)

func main() {

	nomor := 12

	switch nomor {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	default:
		fmt.Println("Tidak ada")
	}
}
