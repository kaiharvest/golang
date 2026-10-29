package main

import "fmt"

type Customers struct {
	name, address string
	age           int
}

func (customers Customers) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customers.name)
}

func main() {
	var customers Customers
	customers.name = "indra"
	customers.address = "indonesia"
	customers.age = 21

	fmt.Println(customers)
	joko := Customers{
		name:    "joko",
		address: "indonesia",
		age:     30,
	}
	fmt.Println(joko)

	budi := Customers{"budi", "indonesia", 17}
	fmt.Println(budi)

	budi.sayHello("agus")
}
