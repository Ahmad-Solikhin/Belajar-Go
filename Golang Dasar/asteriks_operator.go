package main

import "fmt"

/*
Saat membuat variable pointer dan mengubah variable tersebut
Maka variable yang mengacu ke pointer tersebut tidak akan berubah valuenya
Jika ingin mengubah juka semua variable yang mengacu ke variable pointernya bisa menggunakan *
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Tambun"
	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = Address{"Hum", "Ham", "Hem"} Ini gabisa
	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1)
	//fmt.Println(address2)

	*address2 = Address{"Semarang", "Jawa Tengah", "Indonesia"} //Dengan begini address1 juga akan berubah karena digunakan *
	fmt.Println(address1)
	fmt.Println(address2)

}
