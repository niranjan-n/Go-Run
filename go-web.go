package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first Web page")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the first web page by Niranjan N")
}
func view_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey There Fellas!!!!!")
}

func view2_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1> Go is simple</h1>
<p>Learning requires patience!!!!!!
	`)
}
func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("/view", view_handler)
	http.HandleFunc("/view2", view2_handler)
	http.ListenAndServe(":8080", nil)
}
