package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price,string"`
	Discount int    `json:"-"`
	Quantity int    `json:"quality,omitempty"`
	Quality  string `json:"quality_is"`
}

func main() {
	p := Product{
		Name:  "Laptop",
		Price: 50000,
	}
	p_json, err := json.Marshal(p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(p_json))

	//---conmveerting json to go data

	var p_1 Product
	err = json.Unmarshal(p_json, &p_1)
	fmt.Println(p_1)
}
