package main

import (
	"fmt"
)



 type User struct{
	 ID int
	 FisrtName string
	 LastName string
	 Email string
	 IsActive bool
 }

func main() {
	//cara panggil model bikin variabel dan nama modelnya
	user :=User{}

	user.ID=1
	user.FisrtName="Utha"
	user.LastName="ganteng"
	user.Email="a@email.com"
	user.IsActive=true

	user2 :=User{}
	user2.ID=2
	user2.FisrtName="Utha 2"
	user2.LastName="ganteng 2"
	user2.Email="a2@email.com"
	user2.IsActive=false

	fmt.Println(user)
	fmt.Println(user2.FisrtName)
}
