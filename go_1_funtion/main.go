package main

import (
	"fmt"
)

func main() {
	add()
	fmt.Println("==============")
	addWithParameter("1. function dengan parameter")
	fmt.Println("============== untuk function input proses output========")
	gabungan := kalimat("2. klub favorite saya")
	fmt.Println(gabungan)
	fmt.Println("============== 3.untuk function input proses output bag 2========")
	nilai := tambah(90, 10)
	fmt.Println(nilai)
	fmt.Println("============== 4. untuk function calculate input proses dan output nya bisa ambil 2 nilai========")
	//kita buat 2 variabel dulu
	luas, keliling := calculate(10, 2)
	//klo mau ambil luas nya aja kasih tanda _ itu aja
	fmt.Println(luas)
	fmt.Println(keliling)
}

//funtion add
func add() {
	fmt.Println("ini function add")
}

func addWithParameter(kata string) {
	fmt.Println(kata)
}

//funtion ada input parameter(string huruf) proses dan output (string) caranya adalah
func kalimat(huruf string) string {
	//proses
	bahasa := huruf + " " + "juventus"
	return bahasa
}

//function tambah
func tambah(nomor, nomorSatu int) int {
	//angka := nomor + nomorSatu
	return nomor + nomorSatu
}

//function calculate untuk output nya bisa 2
//ouputnya mau luas dan kelilingnya brapa

func calculate(panjang, lebar int) (int, int) {
	luasnya := panjang * lebar
	kelilingnya := 2 * (panjang + lebar)

	return luasnya, kelilingnya
}
