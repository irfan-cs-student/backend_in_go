package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name      string
	Age       int
	Education struct {
		Degree   string
		Semester int
	}
}

func main() {

	json_data := []byte(`{
  "name": "Irfan",
  "age": 22,
  "education": {
    "degree": "BSCS",
    "semester": 5
  }
	}`)

	var Irfan Student
	err := json.Unmarshal(json_data, &Irfan)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("_______Details of a student_____")
	fmt.Println(Irfan.Name)
	fmt.Println(Irfan.Age)
	fmt.Println(Irfan.Education)
	fmt.Println(Irfan.Education.Degree)
	fmt.Println(Irfan.Education.Semester)

	fmt.Println("Details of irfan _____", Irfan)

	//_________other view for triks
	irfan_json := json_data
	fmt.Println("___2nd way of marchal and unmarchal____")
	fmt.Println(irfan_json)
	fmt.Println(string(irfan_json))

	//------convert go data of irfan into json using marchal----

	irfan_json_data, err := json.Marshal(Irfan)
	fmt.Println("______converting irfan data to json using marchal_____")
	fmt.Println(string(irfan_json_data))
	fmt.Println(irfan_json_data)

}
