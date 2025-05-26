package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GitNinja36/mongoApi/router"
)

func main() {
	fmt.Println("MOngo API")
	r := router.Router()
	fmt.Println("Serever is getting started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
