package main

import (
	"net/http"
	"log"
	"encoding/json"
)
type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	book := Book{
		Title: "The Gunsliger",
		Author: "Stephen King",
		Pages: 304,
	}
	json.NewEncoder(w).Encode(book)
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color:red'>Hello Go Server!</h1>"))
}
func main(){
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)
	
	log.Fatal(http.ListenAndServe(":5000", nil))
}
