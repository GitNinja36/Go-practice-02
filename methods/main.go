package main

import "fmt"

func main() {
	rohit := User{"Rohit", "fake@gmail.com", true, 22}
	fmt.Println(rohit)
	// fmt.Printf("rohit details are : %+v", rohit)
	// fmt.Printf("name is : %v \nemail id is : %v", rohit.Name, rohit.Email)
	rohit.GetStatus()
	rohit.newMail()
	fmt.Println(rohit)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}

func (u User) newMail() {
	u.Email = "test2@gmail.com"
	fmt.Println("Email of this user is :", u.Email)
}
