package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a + b
	fmt.Println(c)

	i := 10 // i = 10
	i += 10 // i = i + 10, i = 10 + 10 = 20
	fmt.Println(i)

	i += 5 // i = 20 + 5
	fmt.Println(i)

	j := 10
	j++ // j = j + 1, j = 10 + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
}
