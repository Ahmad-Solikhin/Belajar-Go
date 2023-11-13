package main

import "fmt"

func main() {

	/**
	Digolang bisa menaruh function dalam variable dan nanti bisa digunakan sebagai tipe data
	Ini mirip kayak JS harusnyya
	*/

	goodBye := getGoodBye

	fmt.Println(goodBye("Gayuh"))

}

func getGoodBye(name string) string {
	return "Goodbye " + name
}
