package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type A struct {
	Id int `json:"id"`
}
type Art struct {
	Id       int      `json:"id"`
	Articles Articles `json:"articles"`
}
type Article struct {
	Id      int    `json : "Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

type Articles []Article

var articlesGlobal Articles

func allArticles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","GET")
	if (request.Method == "GET") {
		articles := Articles{
			Article{Id: 1, Title: "Harry Potter and The Philosopher's stone", Desc: " It's a story", Content: "Story"},
			Article{Id: 2, Title: "Harry Potter and The Chamber Of Secrets", Desc: " It's a story", Content: "Story"},
		}
		articlesGlobal = articles
		fmt.Println("Endpoint Hit!!!!!!!!     allArticles")
		json.NewEncoder(writer).Encode(articles)
	}
}

type Arts []Art

var artsGlobal Arts

func all(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","*")
	arts := Arts{
		Art{Id: 1, Articles: articlesGlobal},
		Art{Id: 2, Articles: articlesGlobal},
	}
	artsGlobal = arts
	fmt.Println("Endpoint Hit!!!!!!!!     all")
	json.NewEncoder(writer).Encode(arts)
}
func handleRequests() {
	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/", nothing)
	http.HandleFunc("/post", post)
	http.HandleFunc("/all", all)
	http.HandleFunc("/allArticles", allArticles)
	http.HandleFunc("/a", a)

	//http.HandleFunc("/single/{Id}", single)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func post(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","*")
	if (request.Method == "POST") {
		json.NewEncoder(writer).Encode(artsGlobal)
		fmt.Println("fROM POST")
	}
}
func a(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","*")
	variableA:=A{Id:1}
	fmt.Println("Endpoint Hit!!!!!!!!    variableA ")
	json.NewEncoder(writer).Encode(variableA)
}
func homePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","*")
	fmt.Fprintf(writer, "<h1>Welcome to the HomePage!</h1>")
	fmt.Println("Endpoint Hit: homePage")
}

func nothing(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin","*")
	fmt.Fprintf(writer, "<h1>Wrong hit!!!!!!!!!</h1>")
	fmt.Println("Endpoint Hit: Nothing")
}
func main() {
	handleRequests()
}
