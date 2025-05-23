package main

import "fmt"

func main() {
	// var one *int
	// fmt.Println("value of pointer is ", one)

	myNum := 23
	var ptr = &myNum
	fmt.Println("value of number is : ", ptr)
	fmt.Println("value of number is : ", *ptr)
}
