package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var name byte
	name = 97
	fmt.Println(string(name))

	p1 := []byte("Irfan")
	fmt.Println(p1)

	var nam []byte
	nam = []byte{73, 114, 102, 97, 110}

	fmt.Println(string(nam))

	nam[4] = 109
	fmt.Println(string(nam))

	//---important
	var a byte = 65
	b := []byte("Irfan")
	c := [5]byte{73, 114, 102, 97, 110}

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c[:]))

	d := []byte(`{"name":"irfan","age" :99}`)

	//declaring student variable /stance
	var student Student
	err := json.Unmarshal(d, &student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(student)

}
