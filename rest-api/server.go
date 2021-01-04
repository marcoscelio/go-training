package main

import (
	"api"
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", hello.HelloApi)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
	fmt.Println("Hello Started")
	handleRequests()
}
