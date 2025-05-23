package main

import "fmt"

func main() {
	language := make(map[string]string)

	language["JS"] = "javascript"
	language["RB"] = "ruby"
	language["PY"] = "python"

	//print language
	// fmt.Println("list of all language : ", language)
	// fmt.Println("JS sort for : ", language["JS"])

	//delete
	// delete(language, "RB")
	// fmt.Println("list of all language : ", language)

	//loop
	for _, value := range language {
		fmt.Println("for key v, value is %v\n ", value)
	}
}
