package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "indra"
	names[1] = "dwi"
	names[2] = "prabowo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	values := [...]int{1, 2, 3}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)

}
