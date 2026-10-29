package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{City: "Cepu", Province: "Jawa Tengah", Country: "Indonesia"}
	//address2 := address1 // copy value
	//
	//address2.City = "Blora"
	//fmt.Println(address1) // tidak berubah
	//fmt.Println(address2)

	address1 := Address{City: "Cepu", Province: "Jawa Tengah", Country: "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Blora"
	fmt.Println(address1)
	fmt.Println(address2)
}
