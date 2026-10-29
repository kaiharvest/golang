package main

import "fmt"

func main() {
	nomor := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := nomor[5:9]

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	slice1 := nomor[:5]
	fmt.Println(slice1)

	slice2 := nomor[0:]
	fmt.Println(slice2)

	slice3 := nomor[:]
	fmt.Println(slice3)

	// day
	days := [...]string{"senini", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	dayslice1 := days[5:] // sabtu, minggu
	dayslice1[0] = "sabtu baru"
	dayslice1[1] = "minggu baru"
	fmt.Println(dayslice1) // sabtu baru, minggu baru

	dayslice2 := append(dayslice1, "hari libur") // menambahkan ke dalam days
	dayslice2[0] = "upss"
	fmt.Println(dayslice2) // upss, minggu baru, hari libur
	fmt.Println(days)

	// membuat slice tampa membuat array baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "indra"
	newSlice[1] = "indra"
	// newSlice[2] = "indra" ini error harus nya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "indra")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// array vs slice
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
