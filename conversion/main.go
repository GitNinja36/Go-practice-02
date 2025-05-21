package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to my pizza app :")
	fmt.Println("please rate our pizza between 1 and 5 :")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("thnaks fo rerading :", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Added 1 to your rating :", numRating+1)
	}
}
