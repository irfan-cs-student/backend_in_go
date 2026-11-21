package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Well come to my first go http server !")
}
func main() {

	http.HandleFunc("/", homeHandler)
	fmt.Println("server running at http:/localhost:8080")
	http.ListenAndServe(":8080", nil)

}
