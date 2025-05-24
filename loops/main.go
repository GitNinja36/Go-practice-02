package main

import (
	"fmt"
)

func main() {
	// days := []string{"Sun", "Mon", "Wed", "Fri", "Sat"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// fmt.Printf("index is %v and the value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			// break
			continue
		}
		fmt.Printf("value is %v\n", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("hii there")
}
