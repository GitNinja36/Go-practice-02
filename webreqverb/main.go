package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("status Code : ", response.Status)
	fmt.Println("Content length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count is : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

func PerformJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	    {
		    "name":"Rohit",
		    "college":"NIT",
		    "address":"Durgapur"
	    }
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form Data

	data := url.Values{}
	data.Add("firstName", "rohit")
	data.Add("lastName", "Kumar")

	response, err := http.PostForm(myurl, data)
	checkNilErr(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println(string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
