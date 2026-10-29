package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runingApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runingApp(true)
}
