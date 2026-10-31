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
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	// -------------------------------------------
	// json to strings

	jsonName := []byte(`"irfan"`)
	fmt.Println(jsonName)

	data := jsonData
	var nam string
	fmt.Println(jsonName, data, string(data))
	data[1] = 67

	fmt.Println(jsonName, data, string(data))

	err = json.Unmarshal(jsonName, &nam)

	//json is mutable means can change with index

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(nam)

}
