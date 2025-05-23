package main

import "fmt"

func main() {
	var fruitLists [4]string
	fruitLists[0] = "Apple"
	fruitLists[1] = "Tomato"
	fruitLists[3] = "Peach"
	fmt.Println("fruits are : ", fruitLists)
	fmt.Println("fruits are : ", len(fruitLists))

	var vegiList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegitables : ", vegiList)
}
