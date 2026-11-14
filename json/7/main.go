package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	Company   string
	Employees []struct {
		Name   string
		Age    int
		Skills []string
	}
}

func main() {

	json_of_company := []byte(`{
  "company": "Tech Solutions",
  "employees": [
    {
      "name": "Irfan",
      "age": 22,
      "skills": ["Go", "HTML"]
    },
    {
      "name": "Ali",
      "age": 24,
      "skills": ["Python", "SQL"]
    }
  ]
}`)

	var Employees Emp
	err := json.Unmarshal(json_of_company, &Employees)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(Employees)
	fmt.Println(Employees.Employees[0].Name)

	fmt.Println(Employees.Employees[1].Name)

}
