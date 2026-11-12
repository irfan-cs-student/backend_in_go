// Create a Product struct containing:
// - Name
// - Price
// - Quantity
// Create one product and convert it into JSON using json.Marshal().

// Expected output example:
// {"name":"Laptop","price":85000,"quantity":2}

package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string `json:"pruduct_name"`
	Price    int    `json:"rakam"`
	Quantity int    `json:"kitny_hy"`
}

func main() {

	electronic := Product{
		Name:     "torch",
		Price:    300,
		Quantity: 9,
	}
	data, _ := json.Marshal(electronic)

	fmt.Println(data)
	fmt.Print(string(data))

	json_data := data

	var product_1 Product
	err := json.Unmarshal(json_data, &product_1)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(product_1)

}
