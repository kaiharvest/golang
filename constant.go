package main

import "fmt"

func main() {
	const firstName = "Indra"
	lastname := "Dwi Prabowo"

	// firstName = "Joko" (constant tidak bisa di rubah seperti di javascript)

	fmt.Println(firstName + " " + lastname)

	const (
		fistName1 = "Joko"
		lastName1 = "Gunawan"
	)
	fmt.Println(fistName1)
	fmt.Println(lastName1)
}
