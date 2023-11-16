package main

import "fmt"

/*
Go memiliki function atau keyword new untuk membuat pointer baru
Namun new ini hanya mengembalikan value ke variable kosong, jika variablenya sudah ada isinya maka tidak akan bisa
*/

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := new(Address) //Ini harus berisi tipe datanya tanpa isinya, karena gabisa langsung diinialisasi datanya
	alamat2 := alamat1      //Ini juga akan otomatis jadi pointer ke alamat 1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
