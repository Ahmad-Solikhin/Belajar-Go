package main

import "fmt"

/*
Secara default data yang dikirim ke fucntion di go ada pass by value
jadi variablenya aklan aman dan valuenya akan dicopy
*/

type Address struct {
	City, Country string
}

// Function pass by value atau dicopy valuenya
func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

// Function pass by references
func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	alamat := Address{}
	ChangeCountryToIndonesia(alamat)
	fmt.Println(alamat)

	//Jika variablenya bukan pointer tambahkan & saat mengirim variablenya
	ChangeCountryToIndonesiaPointer(&alamat)
	fmt.Println(alamat)

	alamat2 := new(Address)
	alamat2.Country = "Singapura"
	ChangeCountryToIndonesiaPointer(alamat2)
	fmt.Println(alamat2)
}
