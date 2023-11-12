package main

import "fmt"

func main() {
	//Sama kek tipe data di bahasa pemrograman lainnya sih
	a := 10
	b := 10
	c := 10

	fmt.Println(a + b*c)
	fmt.Println((a + b) * c)

	a += 5
	fmt.Println(a)

	a++
	fmt.Println(a)

	benar := true
	fmt.Println(!benar)
}
