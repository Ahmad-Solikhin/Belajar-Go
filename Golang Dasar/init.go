package main

import (
	"Golang_Dasar/database"
	_ "Golang_Dasar/internal" //Dengan menggunakn _ di awal inport maka init dari packagenya akan tetap diakses tanpa komplain
	"fmt"
)

/*
Ini adalah function yang akan diakses secara otomatis ketika package diakses
Untuk menggunakan ini harus membuat function dengan nama init
kadang kita ingin menjalankan init dalam sebuah package tanpa mengakses function atau field di dalam package tersebut
Hal ini bisa diatasi dengan menggunakan _ sebelum nama package ketika melakukan import
*/

func main() {
	fmt.Println(database.GetConnection())
}
