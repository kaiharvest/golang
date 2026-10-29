package main

import "fmt"

func main() {
	name := "indra"

	switch name {
	case "indra":
		fmt.Println("indra")
	case "joko":
		fmt.Println("joko")
	case "budi":
		fmt.Println("budi")
	default:
		fmt.Println("boleh kenalan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
}
