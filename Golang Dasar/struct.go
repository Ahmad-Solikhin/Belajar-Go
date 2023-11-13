package main

import "fmt"

type Person struct {
	//Kalo di go field di struct diawali dengan kapital ges
	Name        string
	Address     string
	Age, Height int
	Phone       Phone //Ini juga bisa nih masukin struct dalam struct
}

// cara membuat function kepemilikan struct disini adalah dengan cara berikut
func (person Person) sayHello(name string) {
	fmt.Println("Hello", name+", my name is "+person.Name)
}

type Phone struct {
	Number string
}

func main() {
	/**
	strcut merupakan template utnuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	Ini representasi data dalam program yang dibuat
	Data struct akan disimpan dalam field, dan cara penggunaannya menggunakan type kek di typescript
	Ini keknya class deh kalo di java
	Cara menggunakan struct juga harus dibuatb object dulu, cara aksesnya juga sama pake .
	Sama kayak class, struct ini juga bisa naro function
	*/

	var gayuh Person
	fmt.Println(gayuh)

	gayuh.Name = "Gayuh"
	gayuh.Age = 21
	gayuh.Address = "Bekasi"
	gayuh.Height = 173
	gayuh.Phone = Phone{"8618396199"}

	fmt.Println(gayuh)

	amanda := Person{"Amanda", "Bekasi", 22, 169, Phone{"9813628190"}}
	fmt.Println(amanda)

	joko := Person{
		Name:    "Joko",
		Address: "Mars",
		Age:     900,
	}

	fmt.Println(joko)

	jokoPhone := Phone{Number: "7913691"}
	joko.Phone = jokoPhone

	fmt.Println(joko)

	//Ini gabisa dilakukan ges
	//sayHello("Gayuh")

	//Bisanya begini
	gayuh.sayHello("Amanda")
}
