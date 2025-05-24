package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web Request")

	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Println("response is type of :", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
