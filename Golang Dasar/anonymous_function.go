package main

import (
	"fmt"
	"strings"
)

type Blacklist func(string) bool

func main() {

	/**
	Sama aja kayak di js sih harusnya, langsung nulis functionnya di parameternya
	*/

	//Cara pertama
	blaclist := func(name string) bool {
		return strings.EqualFold("anjing", name)
	}
	registerUser("Gayuh", blaclist)

	//Cara kedua
	registerUser("Anjing", func(name string) bool {
		return strings.EqualFold("anjing", name)
	})

}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Save")
	}
}
