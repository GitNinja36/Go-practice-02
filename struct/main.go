package main

import "fmt"

func main() {
	rohit := User{"Rohit", "fake@gmail.com", true, 22}
	// fmt.Println(rohit)
	// fmt.Printf("rohit details are : %+v", rohit)
	fmt.Printf("name is : %v \nemail id is : %v", rohit.Name, rohit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
