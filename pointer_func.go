package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func CountryChagesAddress(address *Address1) {
	address.Country = "Indonesia"
}

func main() {
	address4 := &Address1{} // pointer
	CountryChagesAddress(address4)

	// atau ini juga bisa
	// CountryChagesAddress(&address4)

	fmt.Println(address4)
}
