package main

import (
	"encoding/json"
	"fmt"
)

type Discription struct {
	Brand     string
	Ram       string
	Processor string
}
type product struct {
	Name    string
	Price   int
	Details Discription
}

func main() {

	P_json_data := []byte(`{
  "product": "Laptop",
  "price": 150000,
  "details": {
    "brand": "Lenovo",
    "ram": "8GB",
    "processor": "Intel i7"
  }
}`)
	var mylaptop product
	err := json.Unmarshal(P_json_data, &mylaptop)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mylaptop)

}
