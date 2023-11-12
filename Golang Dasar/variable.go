package main

import (
	"fmt"
)

func main() {
	//Membuat variable dulu tanpa inisialisasi perlu meneyebutkan tipe datanya
	var name string

	name = "Gayuh"
	fmt.Println(name)

	name = "Ahmad"
	fmt.Println(name)

	//Jika valuenya langsung diberikan maka tidak perlu menuliskan tipe datanya
	var secondName = "Solikhin"
	fmt.Println(secondName)

	//Jika langsung inisialisasi variable awal bisa menggunakan cara ini dan tidak perlu menggunakan var serta tipe datanya
	lastName := "Raharjo"
	fmt.Println(lastName)

	//Deklarasi banyak variable sekaligus
	var (
		age        int
		middleName string
		place      = "Bekasi"
	)

	age = 21
	middleName = "Gayuh"

	fmt.Println("My middel name is ", middleName, " with ", age, " years age, and live in ", place)

	//Untuk membuat constant, tidak menggunakan var tapi langsung const. Kek di js ges
	const nik = 2916969

	fmt.Println(nik, age)

	const (
		address = "Bekasi"
		planet  = "Earth"
	)
}
