package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{10, 5, 8, 9, 7}
	total := sum(numbers)
	fmt.Println(total)

	fmt.Println("+++++++++++Quiz2+++++++")
	//+++++++++++Quiz 2 +++++//
	//knpa result, err ini liat dr function calculate yg ouputnya minta int dan error maknnya result, err
	//function calculate(parameternya => 10, 2, "+")
	result, err := calculate(10, 2, "/")
	if err != nil {
		fmt.Println("Telah Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

//kuis pertama disni buat function penjumlahan / sum suatu nilai
//dr sebuah array
//function yg ada paramater namanya numbers untuk inputann
func sum(numbers []int) int {

	//variabel tampung
	var total int
	//proses
	for _, number := range numbers {
		total = total + number
	}
	return total
}

//kuis ke dua ialah ngecek proses dr function calculasi itu berupa penjumlahan, pengurangan , perkalian atau pembagian
//jika tak ada di salah satunya muncul error
//jd funtion yg ada parameter untuk inputan angkanya dan kata2 errornya
//dimna output berupa int (hasil calculate) dan error (pesan error)

//1.berupa inputan => calculate(number, numbers int, info string)
//2. brupa proses switch info {case,default}
//3. berupa output => (int, error)
func calculate(number, numbers int, info string) (int, error) {

	//variabel tampung
	var result int
	var erroResult error
	//disnin kita bisa pakai switch yah
	switch info {
	case "+":
		result = number + numbers
	case "-":
		result = number - numbers
	case "*":
		result = number * numbers
	case "/":
		result = number / numbers
	default:
		erroResult = errors.New("Unknown Operation")
	}
	return result, erroResult
}

//1. input
//2. proses
//3. ouput
