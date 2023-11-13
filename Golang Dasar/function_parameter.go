package main

import (
	"fmt"
	"strings"
)

/*
*
Jika panjang functionnya terlalu panjang baik banyak parameter atau lainnya bisa menggunakan type declaration
Jadi ini kayak membuat alias
*/
type Filter func(string string) string

func main() {
	sayHelloWithFilter("Gayuh", cursedWord)

	filter := cursedWord
	sayHelloWithFilter("anjing", filter)

	sayHelloWithFilter2("anjing", filter)

}

func sayHelloWithFilter(name string, filter func(string2 string) string) {
	fmt.Println("Hello " + filter(name))
}

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

func cursedWord(name string) (returnedName string) {
	returnedName = "..."
	if !strings.EqualFold("anjing", name) {
		returnedName = name
	}

	return
}
