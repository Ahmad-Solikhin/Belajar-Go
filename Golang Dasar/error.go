package main

/*
Go memiliki interface untuk membuat error bernama error dengan bawaan method return Error() string
Untuk membuat error tidak perlu membuat error lagi dari awal
Bisa menggunakan bantuan dari packages errors
*/

import (
	"errors"
	"fmt"
)

func main() {
	result, isError := Pembagian(10, 2)
	if isError == nil {
		fmt.Println("Hasil :", result)
	} else {
		fmt.Println(isError.Error())
	}

	result, isError = Pembagian(10, 0)
	if isError == nil {
		fmt.Println("Hasil :", result)
	} else {
		fmt.Println(isError.Error())
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan 0")
	}

	return nilai / pembagi, nil
}
