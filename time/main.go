package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang:")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.July, 27, 11, 04, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
