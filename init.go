package main

import (
	"fmt"
	"golang_dasar/database"
	_ "golang_dasar/internal"
)

func main() {
	fmt.Println(database.GetConnection())
}
