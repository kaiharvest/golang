package main

import "fmt"

func endApp1() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runingApp1(error bool) {
	defer endApp1()
	if error {
		panic("Error")
	}

}

func main() {
	runingApp1(true)
	fmt.Println("indra dwi")
}
