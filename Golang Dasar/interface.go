package main

import "fmt"

/*
Ini mirip kayak interface di java, jadi ini merupakan kontrak yang nanti harus diimplementasikan
Interface ini dibuat menggunakan type InterfaceName interface { isi method }
*/
type HashName interface {
	GetName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

/*
Untuk implementasi interface tersebut tidak perlu menuliskan secara explisit interface mana yang akan diimplementasikan
Tinggal membuat sesuai dengan kontrak interfacenya
*/
type Person struct {
	Name string
}

// Dengan menambahkan method yang sesuai dengan kontrak HashName ini sudah terhitung sebagai struct yang mengimplementasikan interface Hashname
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	gayuh := Person{Name: "Gayuh"}
	SayHello(gayuh)

	cat := Animal{"Cimo"}
	SayHello(cat)

	fmt.Println(ups1())
	fmt.Println(ups2())
}

/*
Interface kosong
Go ini bukan pahasa pemrograman berorientasi object, kalo di oop biasanya ada Object parent untuk semua tipe data object
Contohnya adalah Object di java
Di go bisa dilakukan dengan membuat interface kosong atau tipe datanya menggunakan any kek di js
kenapa bisa menggunakan any, karena any merupakan type untuk interface kosong
*/
func ups1() any {
	return 1
}

func ups2() interface{} {
	return "Ups"
}
