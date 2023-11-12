package main

import "fmt"

func main() {
	var name = "Gayuh"

	fmt.Println(getHello(name))
	fmt.Print("Hello World")

	name = "Ahmad"
	fmt.Println(getHello(name))
}

func getHello(name string) string {
	return "Hello " + name
}
