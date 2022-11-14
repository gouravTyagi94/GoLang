package main

import "fmt"

func main() {

	//No inheritance in golang; no super or parent
	John := User{"John", "John@go.dev", true, 15}
	
	fmt.Println(John)
	fmt.Printf("Name is %v and Email is %v \n", John.Name ,John.Email)

	//To get more structured details use %+v
	fmt.Printf("John's details are %+v: \n", John)
	
	John.GetStatus()
	John.NewMail()
	fmt.Printf("Name is %v and Email is %v \n", John.Name ,John.Email)
}

//User is public because it needs to be exported out
type User struct {

	//All the fields are public because these needs to be exported out
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "johnty@go.dev";
	fmt.Println("Email of this user is: ", u.Email)
}