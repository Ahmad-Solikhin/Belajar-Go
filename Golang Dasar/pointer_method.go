package main

import "fmt"

/*
Walaupun method menempel di struct, namun data struct yang ada di method adalah pass by value
jadi disarankan method ini menggunakan pointer juga
*/

type Man struct {
	Name string
}

// Ini masih pass by value
func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

// Ini pake pointer
func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	gayuh := Man{"Gayuh"}
	gayuh.Married()
	fmt.Println(gayuh.Name)

	gayuh.MarriedPointer()
	fmt.Println(gayuh.Name)
}
