package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// "math/rand"

func main() {
	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("the sum id is :", myNumberOne+int(myNumberTwo))

	//random number -> using math rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random number -> using crypto rand
	myRandomeNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomeNum)
}
