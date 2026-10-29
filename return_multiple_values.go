package main

import "fmt"

func getNames() (string, string) {
	return "indra", "dwi"

}

func main() {
	// semua data di ambil
	// fistName, lastName := getNames()
	// fmt.Println(fistName, lastName)

	// mengambil data pertama
	fistName, _ := getNames()
	fmt.Println(fistName)
}
