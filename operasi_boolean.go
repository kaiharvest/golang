package main

import "fmt"

func main() {
	nilai := 75
	absen := 80

	nilaiMatkul := nilai >= 80
	nilaiAbsen := absen <= 80

	lulus := nilaiMatkul && nilaiAbsen
	fmt.Println(lulus)
}
