package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	name := "irfan"
	fmt.Println(name)

	jsonData, err := json.Marshal(name)

	if err != nil {
		fmt.Print("error", err)
		return
	}
}
