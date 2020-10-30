package main

import (
	"fmt"
)

func main() {
	//disini kita perumpamaan tbl mahasiswa yg miliki field id dan nama
	//dan tbl mahasiswa bisa bnyak nama dan score

	//nah kan 1 mahasiswa itu 1 map nah klo bnyak mahasiswa > dari 1 map
	//gimnana caranya brati pakai slice untuk menggabungkan suatu map

	//inget yah slice itu []

	//example
	//variabel student dimana satu sclice / [] bisa banyak map
	//dan dimna map nya kan ada key yaitu string dan value nya string
	student := []map[string]string{
		{"name": "utha", "score": "A"},
		{"name": "atho", "score": "A"},
		{"name": "hani", "score": "B"},
		{"name": "mila", "score": "B"},
	}
	//fmt.Println(student)

	//kita cetak satu2
	for _, result := range student {
		fmt.Println(result["name"], "-", result["score"])
	}
}
