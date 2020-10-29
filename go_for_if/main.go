package main

import "fmt"

func main() {
	//for each dan if
	//modulo
	title := "utha ganteng"
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("indexnya", index, "untuk perulangan", string(letter))
		}

	}

	//unutk dapetin jika aiueo panggil
	//pake switch
	for index, letter := range title {
		kata := string(letter)

		switch kata {
		case "a", "i", "u", "e", "o":
			fmt.Println("indexnya", index, "untuk perulangan dan ada pake switch case", string(letter))
		}

	}

	//unutk dapetin jika aiueo panggil
	//pake if
	for index, letter := range title {
		kata := string(letter)

		if kata == "a" || kata == "i" || kata == "u" || kata == "e" || kata == "o" {
			fmt.Println("indexnya", index, "untuk perulangan dan ada pake if", string(letter))
		}

	}
}
