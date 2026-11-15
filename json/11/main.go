package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Age   int `json:"age,omitempty"`
	Email string
}

func main() {

	data := []byte(`{
        "Name": "Irfan",
        "Age": 20,
        "Email": null
    }`)

	var user User

	err := json.Unmarshal(data, &user)

	fmt.Println(err)
	fmt.Println(user)

	var user_json []byte
	user_json, err = json.Marshal(user)
	fmt.Println(err)
	fmt.Println(string(user_json))

}
