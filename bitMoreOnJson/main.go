package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "fake.com", "1234", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "fake1.com", "369369", []string{"full-stack", "js"}},
		{"GO Bootcamp", 99, "fake3.com", "336699", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 199,
		"website": "fake1.com",
		"tags": ["full-stack", "js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//other way
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
