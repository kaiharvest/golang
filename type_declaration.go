package main

import "fmt"

func main() {
	type noKTP string

	nomer := "1111111111"
	fmt.Println(noKTP(nomer))
}
