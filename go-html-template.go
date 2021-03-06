package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsApp struct {
	Title string
	News  string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	n := NewsApp{"News", "Hello there!"}
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, n)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	http.HandleFunc("/", index_handler)
	fmt.Println("Server started...")
	http.ListenAndServe(":8080", nil)
}
