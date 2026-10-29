package main

import "fmt"

func sayHello(fistName string, lastName string) {
	fmt.Println("Hello", fistName, lastName)
}

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	sayHello("indra", "dwi")
	fmt.Println("================")
	result := getHello("indra")
	fmt.Println(result)

	fmt.Println(getHello("joko"))
	fmt.Println(getHello("budi"))
}
