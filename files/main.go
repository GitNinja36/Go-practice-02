package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This need to go in a file "

	file, err := os.Create("./practiceFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./practiceFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Text data inside the file is :", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
