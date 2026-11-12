// Given this JSON:
// {
//     "name": "Irfan",
//     "age": 22,
//     "skills": ["Go", "HTML", "CSS"]
// }

// 1. Create a suitable struct.
// 2. Use json.Unmarshal().
// 3. Print the name.
// 4. Print the age.
// 5. Print all skills.

package main

import (
	"encoding/json"
	"fmt"
)

type Fellow struct {
	Name   string
	Age    int
	Skills []string
}

func main() {

	json_data := []byte(`{
    "name": "Irfan",
    "age": 22,
    "skills": ["Go", "HTML", "CSS"]

      }`)

	fmt.Println("json data___", string(json_data))

	//-------json into proper go data in using Studnet struct

	var fellow Fellow
	err := json.Unmarshal(json_data, &fellow)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fellow)
}
