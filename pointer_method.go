package main

import "fmt"

type Man struct {
	Name string
}

// untuk method sangat disarankan menggunakan pointer
func (m *Man) Married() {
	m.Name = "Mr. " + m.Name
}

func main() {
	names := Man{"indra"}
	names.Married()

	fmt.Println(names.Name)
}
