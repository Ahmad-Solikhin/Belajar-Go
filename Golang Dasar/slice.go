package main

import "fmt"

func main() {
	/**
	Slice ini kayak array, tapi panjangnya bisa dinamis
	Slice dan array berhubungan karena slice adalah data yang mengakses sebagian atau seluruh data di array
	Keknya ini mirip List di java
	Slice memiliki 3 informasi penting:
	 	pointer : penunjuk data pertama di array
		length : panjang slice
		capicity : kapasitas dari slice
	Membuat slice dari array yang sudah ada, mirip kek di python
		array[low:high] : membuat slice dari array mulai dari index low samapai index high
		array[low:] : membuat slice dari array index low sampai paling terakhir
		array[:high] : mmebuat slice dari awal sampai high yang ditentukan
		array[:] : membuat slice dari seluruh data di array
		low itu menentukan posisi index data mulai diambil, dan highnya adalah posisi data menggunakan index + 1
	*/

	names := [...]string{
		"Ahmad", "Gayuh", "Amanda", "Joko", "Budi", "Rahmat",
	}

	slice1 := names[1:5]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[:]
	fmt.Println(slice3)

	/**
	Function dari slice
		- len : mendapatkan panjang dari slice
		- cap : mendapatkan kapasitas dari slice
		- append(slice, data) : membuat slice baru dengan mendambahkan data ke posisi terakhir dari slice, jika posisinya penuh maka akan dibuatkan array baru
		- make([]DataType, length, capacity) : membuat slice baru dari array baru
		- copy(destination, src) : copy slice lama ke slice baru
	*/
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	//Ini akan mengubah data di array yang digunakan untuk membaut slice, karena slice ini masih berada dalam kapasitas dan tidak dibuatkan array baru
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)
	fmt.Println(cap(daySlice1))

	//Slice ini akan membuat array baru, karena array sebelumnya yang digunakan sudah melebihi kapasitas dari array yang digunakan sebelumnya
	//Oleh karena itu, modifikasi dalam slice ini tidak akan mempengaruhi array awal yang digunakan dalam membuat slice
	daySlice2 := append(daySlice1, "hari baru")
	fmt.Println(cap(daySlice2))
	daySlice2 = append(daySlice2, "Anjayy")
	fmt.Println(cap(daySlice2))
	daySlice2[0] = "Upss"
	fmt.Println(daySlice2)
	fmt.Println(days)

	//Membuat slice tanpa array yang sudah ada
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Gayuh" //Best practice menambahkan data ke slice tidak menggunakan index, tapi di append
	newSlice[1] = "Ahmad"
	newSlice = append(newSlice, "Joko")
	fmt.Println()
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//Copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	//Jika dilakukan copy, maka destinasi slicenya ini tentu akan menggunakan array baru miliknya, jadi perubahan pada slice ini tidak akan mengubah atau mempengruhi slice atau array sebelumnya
	copy(toSlice, fromSlice)

	toSlice[1] = "Sebbbb"

	fmt.Println()
	fmt.Println()
	fmt.Println(fromSlice)
	fmt.Println(days)
	fmt.Println(toSlice)

	/**
	hati-hati dalam membuat array, pastikan saat membuat array yang terbentuk adalah benar2x array bukan slice begitu juga sebaliknya
	*/

	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}
	fmt.Println()
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
