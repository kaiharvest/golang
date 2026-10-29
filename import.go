package main

import (
	"fmt"
	"golang_dasar/halper"
)

func main() {
	result := halper.SayHello("indra")
	fmt.Println(result)

	fmt.Println(halper.Application)
}
