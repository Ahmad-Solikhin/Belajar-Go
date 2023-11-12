package main

import "fmt"

func main() {
	//Digunakan untuk memberikan alias untuk tipe data sebelumnya
	//Contoh

	type NoKTP string

	var ktpGayuh NoKTP = "78136913"
	fmt.Println(ktpGayuh)
	fmt.Println(NoKTP("710937190"))

	stringKtp := string(ktpGayuh)
	fmt.Println(stringKtp)
}
