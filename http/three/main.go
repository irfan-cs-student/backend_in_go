package main

import (
	"fmt"
	"net/http"
)

// GET /student
func studentHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request received!")
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.Path)

	fmt.Fprintln(w, "Student Information")
	fmt.Fprintln(w, "Name: Irfan")
	fmt.Fprintln(w, "Course: Go Backend")
	fmt.Fprintln(w, "Semester: 5")
}

// GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Student Management System")
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/student", studentHandler)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
