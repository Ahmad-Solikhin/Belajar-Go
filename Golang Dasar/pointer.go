package main

import "fmt"

/*
Semua variable di golang adalah pasing by value bukan references
Jadi saat mengirim variable kedalam function maka datanya akan diduplikat terlebih dahulu
Jadi data yang sebelumnya akan aman jika diubah

Nah jadi pointer ini digunakan ketika ingin melakukan reference bukan mengcopy datanya
Dengan menggunakan pointer ini maka melakukan pass by reference bukan pass by value
Karena jika melakukan pass by value terus menerus maka akan boros memory
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	{
		//Ini pass by value
		address2 := address1 //Saat di assign ke address2 ini dicopy value address1 nya

		address2.City = "Karawang" //Ini gabakal ngubah value di address1, karena value di address2 ini merupakan copy dari address1 bukan reference

		fmt.Println(address1)
		fmt.Println(address2)
		fmt.Println("=========================")
	}

	{
		//Ini pass by reference
		/*
			Untuk membuat variable pointer ke variable lain bisa menggunakan & di awal variable yang ingin dilakukan reference
			Jadi karena reference ketika melakukan perubahan di variable yang di pass by reference maka value di variable sebelumnya akan berubah juga
		*/
		var address3 *Address = &address1
		address3.Country = "UA"
		fmt.Println(address1)
		fmt.Println(address3)

	}
}
