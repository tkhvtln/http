package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page!")
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about", about_page)

	fmt.Printf("Starting server at :8080")
	http.ListenAndServe(":8080", nil)

}

func main() {
	handleRequest()
}
