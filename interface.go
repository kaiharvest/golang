package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello1(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	sayHello1(Person{Name: "Indra"})

	animal := Animal{"Gajah"}
	sayHello1(animal)
}
