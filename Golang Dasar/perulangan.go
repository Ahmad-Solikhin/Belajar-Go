package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Ini juga sama kayak perulangan di bahasa pemrograman lainnya

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-" + strconv.Itoa(counter))
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Value ke-" + strconv.Itoa(i))
	}

	//DI go bisa melakukan iterasi dari semua data dalam collection menggunakan for range
	names := []string{"Ahmad", "Solikhin", "Gayuh", "Raharjo"}
	for index, name := range names {
		fmt.Println("Index ke-" + strconv.Itoa(index) + " : " + name)
	}

	//Karena di go semua variable harus dipake, jika tidak membutuhkan index nya bisa diubah menjadi "_"
	for _, name := range names {
		fmt.Println("Value : " + name)
	}

}
