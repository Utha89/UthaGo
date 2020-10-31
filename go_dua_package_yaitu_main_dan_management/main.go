package main

import (
	"fmt"
	"go_dua_package_yaitu_main_dan_management/management"
)
func main()  {
	//user :=User{1,"utha","ganteng","a@mail.com",true}
	user :=management.User{1,"utha","ganteng","a@mail.com",true}

	result :=user.GetData()
	fmt.Println(result)

	user2 :=management.User{1,"utha dybala","ganteng","a@mail.com",true}
	fmt.Println(user2.GetData())

	usersAll :=[]management.User{user, user2}
	group :=management.Group{"Gamer", user, usersAll, true}
	group.GetDataAllMethod()
	
}