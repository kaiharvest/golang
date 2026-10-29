package main

import "fmt"

func getCompleteName() (firstName, midName, lastName string) {
	firstName = "indra"
	midName = "dwi"
	lastName = "prabowo"

	return firstName, midName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
