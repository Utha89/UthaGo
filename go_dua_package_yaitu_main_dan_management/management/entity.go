package management
import "fmt"
type User struct{
	ID int
	FisrtName string
	LastName string
	Email string
	IsActive bool
}

func(user User) GetData() string {
	return fmt.Sprintf("Name : %s %s, Email : %s ", user.FisrtName, user.LastName, user.Email)
	
}

type Group struct{
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func(group Group) GetDataAllMethod()  {
	fmt.Printf("ini struct function pakai method, Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	// kan td di gruop modelnya slice / map yak yg property nya 
	//Users yaitu kumpulan [] dr user maka kita bisa di for each
	 for _,hasil:=range group.Users{
	
		 fmt.Println("")
		 fmt.Printf(hasil.FisrtName)
	 }

	
	//next ke main lg panggil si function getDataAll
}
//funct untuk model yg dalemnya ada model lg
func GetDataAll(group Group)  {
	fmt.Printf("ini struct function pakai parameter,Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	// kan td di gruop modelnya slice / map yak yg property nya 
	//Users yaitu kumpulan [] dr user maka kita bisa di for each
	 for _,hasil:=range group.Users{
	
		 fmt.Println("")
		 fmt.Printf(hasil.FisrtName)
	 }

	
	//next ke main lg panggil si function getDataAll
}
//Disini kita panggil model / struct dr suatu function

func getData(user User) string {
	result :=fmt.Sprintf("Name : %s %s, Email : %s ", user.FisrtName, user.LastName, user.Email)
	return result
}

