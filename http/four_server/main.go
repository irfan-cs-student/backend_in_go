// Create a Go HTTP server that provides information about a library.
package main

import (
	"fmt"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("wellcome to home !")
	fmt.Fprintln(w, "Welcome to My Library")
}
func bookshandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("wellcome to books sections")

	fmt.Fprintln(w, "Available Books: Go, Python, C++")

}
func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("________________about _________")

	fmt.Fprintln(w, "PICK THE right book that is suitable for you")

}
func fictionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("______films sections__")
	fmt.Fprintln(w, "u got films sections")

}

func main() {

	http.HandleFunc("/", homehandler)
	http.HandleFunc("/books", bookshandler)
	http.HandleFunc("/about", abouthandler)
	http.HandleFunc("/books/fiction", fictionHandler)

	http.ListenAndServe(":8080", nil)

}
