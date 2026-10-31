package main

import "fmt"

// byte is an alias for unit8(form 0 to 256, note 256 cause eroors)

func main() {
	a := 2
	var b byte = 5
	// var aa byte=999   error beacuse values is bigger than 255

	fmt.Print(a, b)

	name := "irfan"
	fmt.Println(name)

	//string to byte
	byteName := []byte(name)
	fmt.Println(byteName)
	fmt.Println(string(byteName), "----------------")

	//int to charaters
	n := 99
	fmt.Printf("%c\n", n)
	fmt.Println(string(n))

	//byte to strings
	nameINnu := []byte(name)
	irfanBYte := byteName
	fmt.Println("irfan in byte___", nameINnu)
	fmt.Println("irfan in byte___", irfanBYte)

	//
	abc := []byte("ع")
	abcd := []rune("ع")
	fmt.Println(abc)
	fmt.Println(abcd)

	namm := "علی"

	letters := []rune(namm)
	letters_1 := []byte(namm)

	fmt.Println(len(letters))
	fmt.Println(len(letters_1))

	numbers := []int{10, 20, 30}

	names := []string{"Ali", "Irfan"}

	prices := []float64{10.5, 20.75}

	data := []byte{65, 66, 67}

	values := []uint8{10, 200, 255}

	fmt.Println(numbers)
	fmt.Println(names)
	fmt.Println(prices)
	fmt.Println(data)
	fmt.Println(values)

	mixed := []any{10, "Irfan", 20.5, true}
	fmt.Println(mixed)

}
