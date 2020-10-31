package main

import(
	"fmt"
)

type User struct{
	ID int
	FisrtName string
	LastName string
	Email string
	IsActive bool
}
func main()  {
	user :=User{ID:1,FisrtName:"Utha",LastName:"Ganteng",Email:"a@yahuu.com",IsActive:true}

	result :=user.getData()
	fmt.Println(result)
}
func(user User) getData() string {
	return fmt.Sprintf("Name : %s %s, Email : %s ", user.FisrtName, user.LastName, user.Email)
	
}