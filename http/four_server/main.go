// Create a Go HTTP server that provides information about a library.
package main

import "net/http"

func main() {

	http.HandleFunc("/", homehandler)
	http.HandleFunc("/books", bookshandler)
	http.HandleFunc("/about", abouthandler)

	http.ListenAndServe(":8080", nil)

}
