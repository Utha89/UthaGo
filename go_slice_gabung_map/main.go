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

	fmt.Println("+++++++++++++++++++++++Kuiz 1 mencari nilai rata2 rumusnya kan total data /sum(data)++++++++++")

	//kuiss
	//1. mencari nilai rata2 rumusnya kan total data /sum(data)
	//total datanya = 8
	//sum(100+80+75 dst)

	//ini array
	scores := [...]int{100, 80, 75, 92, 70, 93, 88, 67}
	length := len(scores)

	//suatu variabel
	var tampung int

	//ini for each
	for _, result := range scores {
		tampung = tampung + result
	}

	average := float64(tampung) / float64(length)
	fmt.Println(average)

	fmt.Println("+++++++++++++++++++++++Kuiz 2++++++++++")

	//kuiss 2
	//memasukkan suatu data > 90 kedalem variabel

	//ini array
	scores2 := [...]int{100, 80, 75, 92, 70, 93, 88, 67}
	//ini slice
	var goodScores []int

	//ini for each
	for _, data := range scores2 {
		if data >= 90 {
			//bikin slice baru
			goodScores = append(goodScores, data)
		}

	}

	fmt.Println(goodScores)
}
