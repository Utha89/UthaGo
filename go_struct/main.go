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

//model yg didalem nya ada model lg
type Group struct{
		Name string
		Admin User
		Users []User
		IsAvailable bool
}

func main() {
	//cara panggil model bikin variabel dan nama modelnya
	user :=User{ID:1,FisrtName:"Utha",LastName:"Ganteng",Email:"a@yahuu.com",IsActive:true}

	// user.ID=1
	// user.FisrtName="Utha"
	// user.LastName="ganteng"
	// user.Email="a@email.com"
	// user.IsActive=true

	user2 :=User{}
	user2.ID=2
	user2.FisrtName="Utha 2"
	user2.LastName="ganteng 2"
	user2.Email="a2@email.com"
	user2.IsActive=false

	fmt.Println(user)
	fmt.Println(user2.FisrtName)

	fmt.Println("+++++++++++++++Method yg panggil model++++++++++++")
	data :=getData(user)
	data2 :=getData(user2)
	fmt.Println(data)
	fmt.Println(data2)

	fmt.Println("+++++++++++++++model yg dalemnya ada model lg++++++++++++")
	//main untuk panggil model yg dalemnya ada model lg
	//Group ada nama modelnya yah
	//untuk Name isi gamer
	//untuk Admin krena dia acuannya model di User maka kita bisa isi yg udah dibuat di main yg panggil model User
	//yaitu user
	//untuk Users artinya kumulan dr beberapa user dr model User terus ada slice[]
	//nah karena dr beberapa user dr model User maka kita buat variabelnya dulu nampung
	 
	//gini caranya masukin user dan user2
	usersAll :=[]User{user, user2}
	group :=Group{"Gamer", user, usersAll, true}

	//Pangil function getDataAll
	getDataAll(group)

}


//Disini kita panggil model / struct dr suatu function

func getData(user User) string {
	result :=fmt.Sprintf("Name : %s %s, Email : %s ", user.FisrtName, user.LastName, user.Email)
	return result
}

//funct untuk model yg dalemnya ada model lg
func getDataAll(group Group)  {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	// kan td di gruop modelnya slice / map yak yg property nya 
	//Users yaitu kumpulan [] dr user maka kita bisa di for each
	 for _,hasil:=range group.Users{
		 fmt.Println("User By Group :")
		 fmt.Printf(hasil.FisrtName)
	 }

	
	//next ke main lg panggil si function getDataAll
}