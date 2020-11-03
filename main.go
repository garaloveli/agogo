package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	ID			string `json:id`
	Title 	string `json:"Title"`
	Desc 		string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article = []Article {
	Article { ID:"1", Title:"Test Article", Desc:"Test Description", Content:"Hello World" },
	Article { ID:"2", Title:"Test Article 2", Desc:"Test Description 2", Content:"Hello World 2" },
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func setArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post article hit")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	// fmt.Fprintf(w, "Key: " + key)
	for _, article := range articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	// creates a new instances of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", setArticles).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API Test Sample")
	handleRequest()
}
