package main

import "fmt"

func random() interface{} {
	return "ok"
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	// jangan sampai konversi ke int
	// resultInt := result.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
