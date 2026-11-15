package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"user_name"`
	Password string `json:"-"` //completely ignore the fields
}

// This ("-") is useful for hiding sensitive fields from API responses.
func main() {

	user_1 := User{
		Name:     "Irfan",
		Password: "12345",
	}

	fmt.Println(user_1)

	//------converting to json

	user_1_json, err := json.Marshal(user_1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("___user_1 data to json______")
	fmt.Println(string(user_1_json))

	//converting json to user_1 go data
	fmt.Println("___user_1 json data into go data ______")

	var Orignal_user_1 User

	err = json.Unmarshal(user_1_json, &Orignal_user_1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(Orignal_user_1)

}
