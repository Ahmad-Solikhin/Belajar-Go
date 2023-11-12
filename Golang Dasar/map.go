package main

import "fmt"

func main() {

	/**
	Pada array atay slice untuk mengakses datanya harus menggunakan indexnya yang berupa angka dimulai dari 0
	namun pada map bisa menggunakan key-value yang bisa ditentukan tipe data dari keynya
	Ini sama aja kayak map di java deh harusnya, yang penting keynya harus unique
	Untuk map panjangnya juga bebas, bisa sebanyak2xnya
	*/

	var person1 map[string]string = map[string]string{}
	person1["name"] = "Gayuh"
	person1["age"] = "21"
	person1["address"] = "Bekasi"

	fmt.Println(person1)
	fmt.Println(person1["name"])

	person2 := map[string]string{
		"name":    "Amanda",
		"age":     "21",
		"address": "Bekasi",
	}

	fmt.Println(person2)
	fmt.Println("Person value : ", person2["none"])

	/**
	Function yang biasa digunakan di map
	-	len(map) : mendapatkan jumlah data di map
	-	map[key] : mengambil data di map dengan key
	-	map[key] = value : Mengubah data di map dengan keynya
	- 	make(map[TypeKey]TypeValue) : Membuat map baru
	-	delete(map, key) : menghapus data di map dengan keynya
	*/

	book := make(map[string]string) //map[string]string{}
	book["title"] = "Buku ABC"
	book["author"] = "Gayuh"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
