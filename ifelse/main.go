package main

import "fmt"

func main() {
	logInCount := 10
	var msg string

	if logInCount < 10 {
		msg = "Regular user"
	} else if logInCount > 10 {
		msg = "watch out"
	} else {
		msg = "something else"
	}
	fmt.Println("msg : ", msg)

	//extra
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}
}
