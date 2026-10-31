package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade string
}

func main() {

	// Example 1: Working with bytes

	data := []byte{72, 101, 108, 108, 111}

	fmt.Println("Bytes:", data)
	fmt.Println("String:", string(data))

	message := "Hello Irfan"
	bytes := []byte(message)

	fmt.Println("Original:", message)
	fmt.Println("Bytes:", bytes)
	fmt.Println("Back to string:", string(bytes))

	// Example 2: JSON as bytes

	jsonData := []byte(`{"Name":"Irfan","Age":22,"Grade":"A"}`)

	fmt.Println("\nJSON bytes:", jsonData)
	fmt.Println("JSON string:", string(jsonData))

	// Example 3: JSON to Go struct

	var student Student

	err := json.Unmarshal(jsonData, &student)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nStudent details:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Grade:", student.Grade)

	// Example 4: Go struct to JSON

	newJSON, err := json.Marshal(student)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nConverted JSON:", string(newJSON))
}
