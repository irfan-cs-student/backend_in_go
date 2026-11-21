package main

import (
	"fmt"
	"net/http"
)

// Route 1: Home
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Go backend!")
}

// Route 2: About
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is Irfan.")
	fmt.Fprintln(w, "I am learning Go backend development.")
}

// Route 3: Skills
func skillsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My Skills:")
	fmt.Fprintln(w, "HTML")
	fmt.Fprintln(w, "CSS")
	fmt.Fprintln(w, "C++")
	fmt.Fprintln(w, "Go")
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/skills", skillsHandler)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
