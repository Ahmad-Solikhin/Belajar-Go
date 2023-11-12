package main

import (
	"fmt"
)

func main() {
	/**
	Ini juga sama kayak switch case di bahasa pemograman lainnya
	*/

	var name string
	name = "Amanda"

	switch name {
	case "Amanda":
		fmt.Println("Hello Amanda")
	case "Gayuh":
		fmt.Println("Hello Gayuh")
	default:
		fmt.Println("Lu siapa?")

	}

	//Bisa juga menggunakan short term di dalam switch
	switch length := len(name); length > 5 {
	case true:
		println("Sesuai")
	case false:
		println("Kependekan")
	}

	//Switch tanpa kondisi, jika tidak mau bisa menambahkan kondisi tersebut pada setiap case nya
	switch {
	case len(name) > 10:
		println("Kepanjangan")
	case len(name) > 5:
		println("Cukup")
	default:
		println("Kependekan")

	}
}
