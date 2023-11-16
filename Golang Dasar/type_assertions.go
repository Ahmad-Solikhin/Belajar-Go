package main

import (
	"fmt"
)

func random() any {
	return 100
}

func main() {
	/*
		Type assertions merupakan mlakukan casting tipe data any menjadi tipe data lain
		Jadi ini sering di temui di tipe data interface kosong
		Caranya adalah dengan menggunakan variable_any.(tipe data)
		Jika ada kesalahan dalam melakukan konversi maka akan terpicu panic
	*/

	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//resultInt := result.(int)
	//fmt.Println(resultInt)

	/*
		Ada cara yang lebih aman dalam melakukan konversi tipe data, hal ini adalah menggunakan switc case
		Ini kayak fiture baru di java 21 itu ges
	*/

	switch result.(type) {
	case string:
		fmt.Println("String :", result)
	case int:
		fmt.Println("Int :", result)
	default:
		fmt.Println("Unknown")
	}

}
