package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("indra", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
