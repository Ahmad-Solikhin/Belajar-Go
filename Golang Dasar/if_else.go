package main

import "fmt"

func main() {

	/**
	Ini sama kayak bahasa pemrograman lainnya, buat bandingin string langsung pake == aja di go
	*/

	name := "Amanda"

	if name == "Gayuh" {
		fmt.Println("Benar")
	} else if name == "Amanda" {
		fmt.Println("Panda")
	} else {
		fmt.Println("Salah")
	}

	//Menggunakan if short statement, jadi bisa deklarasi variable dulu langsung di parameter if nya, ini keknya cuma ada di golang
	if length := len(name); length > 2 {
		fmt.Println("Betul")
	}

}
