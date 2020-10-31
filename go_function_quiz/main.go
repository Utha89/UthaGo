package main

import "fmt"

func main() {
	numbers := []int{10, 5, 8, 9, 7}
	total := sum(numbers)
	fmt.Println(total)
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

//1. input
//2. proses
//3. ouput
