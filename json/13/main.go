package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Age *int `json:"age"`
}

func main() {

	x := 10

	s := Student{
		Age: &x, //age pointing to x value 10
	}

	fmt.Println(*s.Age)       //10
	fmt.Println(s.Age == nil) //flase

	data1 := []byte(`{"age": 25}`)
	json.Unmarshal(data1, &s) //age gets 25

	fmt.Println(*s.Age)       //25
	fmt.Println(s.Age == nil) //false

	data2 := []byte(`{"age": null}`)
	json.Unmarshal(data2, &s)

	fmt.Println(s.Age == nil) //s student updates data by getting from data2,so nill

	data3 := []byte(`{}`) //age here is not provided so unmarchal doest cahnge its prevoius values
	json.Unmarshal(data3, &s)

	fmt.Println(s.Age == nil) //true
}
