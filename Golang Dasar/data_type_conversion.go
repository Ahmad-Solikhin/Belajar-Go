package main

import "fmt"

func main() {
	//Konversi tipe data int
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	//Jika dikonversi menjadi nilai yang lebih kecil jangkaunnya maka akan menyebabkan masalah dimana dia akan mengulang dari awal kembali
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//Konversi tipe data string
	var name = "Gayuh"
	var e = name[0]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
