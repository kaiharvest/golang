package main

import (
	"fmt"
)

func main() {
	name := "indra"

	if name == "indra" {
		fmt.Println("Hello " + name)
	} else if name == "joko" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Hello boleh kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
