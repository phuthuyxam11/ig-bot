package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authentication service!")
	fmt.Println("Endpoint Hit: Authentication service")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	handleRequests()
}
