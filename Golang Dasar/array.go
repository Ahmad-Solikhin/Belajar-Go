package main

import "fmt"

func main() {
	//Ini array yang biasa, jadi harus nentuin jumlah datanya di awal dan tidak bisa bertambah kecuali buat data baru
	var names [2]string

	names[0] = "Gayuh"
	names[1] = "Ahmad"

	fmt.Println(names[0])
	fmt.Println(names)

	//Membuat array secara langsung
	values := [3]int{10, 20, 30}

	fmt.Println(values)

	//Function yang bisa digunakan pada array
	/**
	- len : mendapatkan panjang array
	- array[index] : mendapatkan value
	- array[index] = values : assign data ke array berdasarkan indexnya
	- [...] : menentukan jumlah awal sesuai data yang di assign
	*/
	//Dalam golang tidak bisa menghapus data di dalam array, karena saat pertama dibuat semua datanya sudah di assign dengan value 0 atau ""

	var addresses = [...]string{
		"Bekasi", "Jakarta", "Semarang",
	}

	fmt.Println(len(addresses))

	var news [2]string
	fmt.Println(news)
	fmt.Println(len(news))
}
