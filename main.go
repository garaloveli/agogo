package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoiint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}