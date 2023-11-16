package main

import (
	"Golang_Dasar/helper"
	"fmt"
)

/*
Di golang untuk mnentukan access modifier dengan menentukan namanya saya
Kalo diawali dengan kapital merupakan public
Kalo huruf kecil artinya protected
*/

func main() {
	result := helper.SayHello("Gayuh")
	fmt.Println(result)
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) Ini tidak bisa dibuka karena protected
	//fmt.Println(helper.sayGoodBye("Gayuh"))
}
