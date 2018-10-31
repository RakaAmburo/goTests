package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this dimentions!")
	fmt.Println("Endpoint Hit: homePage")

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func main() {
	handleRequests()
}
