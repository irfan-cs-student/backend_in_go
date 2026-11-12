package main

import "fmt"

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

}
