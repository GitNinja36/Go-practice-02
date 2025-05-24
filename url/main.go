package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com:8080/shop?item=book&price=299&currency=INR"

func main() {
	// fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	checkNilErr(err)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Println("The type of query params are :", qParams)

	fmt.Println(qParams["currency"])

	for _, val := range qParams {
		fmt.Println("params is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/shop",
		RawQuery: "user=rohit",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
