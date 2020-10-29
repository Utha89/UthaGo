package main

import (
	"fmt"
)

func main() {

	var age int = 20
	//atau
	angka := 20

	fmt.Println(age)
	fmt.Println(angka)

	//belajar array cara 1
	//ada 5 jumlah array nya dan kkembaliannya string
	var kata [5]string

	kata[0] = "flutter"
	kata[1] = "PHP"
	kata[2] = "Go"
	kata[3] = "Dart"
	kata[4] = "Java"

	//belajar array cara 2
	//ada 5 jumlah array nya dan kkembaliannya string
	bahasa := [5]string{"Flutter", "PHP", "Go", "Dart", "Java"}

	fmt.Println(kata)
	fmt.Println(bahasa)
	fmt.Println("______________________")
	//Untuk mengetahui panjang array
	panjangKata := len(bahasa)
	fmt.Println(panjangKata)
	fmt.Println("______________________")
	//Just info untuk hitung jumlah array atau nentuin jumlah array kan yg udah2
	//[5] di go bisa [...]
	//nah klo nulis array vertikal terakhirnya tetep pake , klo horizontal tnpa koma
	//examplenya

	kalimat := [...]string{
		"Flutter",
		"PHP",
		"Go",
		"Dart",
		"Java",
	}
	fmt.Println("trik nampilin array model vertikal")
	fmt.Println("kalimatnya:", kalimat)
	//nah klo di php lnjutin dr string pake . contoh : $utha."saya"
	//nah klo di flutter(dart) dan JS lnjutin dr string pake + contoh : "saya" + utha
	//nah klo di Golang lnjutin dr string pake , contoh : "saya", utha

	//Nampilin array lalu kita for each
	fmt.Println("______________________")
	sepakBola := [...]string{
		"JUVE",
		"BARCA",
		"Timnas u19",
	}

	fmt.Println("for each yg ambil data dr array :")
	for index, kal := range sepakBola {

		fmt.Println("indexnya", index, "Club nya", kal)
	}
}
