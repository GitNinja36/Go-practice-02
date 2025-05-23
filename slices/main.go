package main

import (
	"fmt"
	"sort"
)

func main() {
	// var fruitLists = []string{"Apple", "Tomato", "Peach"}
	// fmt.Println("fruit list : ", fruitLists)

	// //adding new value in the sclices
	// fruitLists = append(fruitLists, "lichi")
	// fmt.Println("fruit list : ", fruitLists)

	// //slceing
	// fruitLists = append(fruitLists[1:])
	// fmt.Println("fruit list : ", fruitLists)

	highScore := make([]int, 4)
	highScore[0] = 396
	highScore[1] = 369
	highScore[2] = 639
	highScore[3] = 963
	// highScore[4] = 936

	highScore = append(highScore, 333, 666, 999)
	sort.Ints(highScore)
	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove value from slices based on index
	var cources = []string{"javascript", "recatJs", "Swift", "python", "ruby"}
	fmt.Println("cources are : ", cources)
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println("cources are : ", cources)
}
