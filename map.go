package main

import "fmt"

func main() {
	person := map[string]string{
		"first": "Indra",
		"last":  "Dwi",
	}

	fmt.Println(person)
	fmt.Println(person["first"])
	fmt.Println(person["last"])

	names := map[int]string{
		1: "Nugroho",
		2: "Joko",
	}

	fmt.Println(names)
	fmt.Println(names[1])
	fmt.Println(names[2])

	books := map[string]string{
		"autor": "Izza",
		"judul": "lalalalalala",
	}

	delete(books, "judul")
	fmt.Println(books)
}
