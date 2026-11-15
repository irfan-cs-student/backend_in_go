package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   *int
	Email string
}

func main() {

	data := []byte(`{
        "Name": null,
        "Age": 20,
        "Email": null
    }`)

	var s Student

	json.Unmarshal(data, &s)

	fmt.Println(s.Name)
	fmt.Println(s.Age == nil)
	fmt.Println(*s.Age)
	fmt.Println(s.Email)
}
