package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	names := []string{"indra", "dwi", "prabowo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// menggunakan range buat array, slice, dan map
	names1 := []string{"indra", "dwi", "prabowo"}
	for index, name := range names1 {
		fmt.Println("index", index, "=", name)
	}

	// map
	names2 := map[string]string{
		"nama1": "indra",
		"nama2": "dwi",
		"nama3": "prabowo",
	}

	for key, name := range names2 {
		fmt.Println("key", key, "=", name)
	}

	// mengecualian untuk mengambil key
	for _, name := range names {
		fmt.Println(name)
	}

}
