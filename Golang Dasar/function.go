package main

import (
	"fmt"
)

func main() {

	/**
	Membuat function di golang menggunakan kata kucni func dan memberikan nama functionnya dan diberikan parameter jika ingin

	*/

	helloWord()

	sayHello("Gayuh")
	sayHello("Amanda")
	sayHelloTo("Gayuh", "Amanda")

	fmt.Println(gettingHello("Gayuh"))

	var result = gettingHello("Amanda")
	fmt.Println(result)

	_, multipleResult, _, _ := getFullName() //jadi misal gamau semua data yang dipake bisa pake _ ges
	fmt.Println(multipleResult)

	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println()
	println(sumAll(10, 20, 30, 40))

	//Jika kalau ada kasus functionnya sudah menggunakan variadic namun data yang dimiliki sudah dalam bentuk slice
	//Hal ini masih bisa menggunakan slice menjadi variadic
	numbers := []int{10, 20, 30}
	fmt.Println(sumAll(numbers...))

}

// Function return multiple value
// Caranya adalah dengan menuliskan semua kembalian valuenya setelah function menggunakan ()
func getFullName() (string, string, string, string) {
	return "Ahmad", "Solikhin", "Gayuh", "Raharjo"
}

// Function return value
func gettingHello(name string) string {
	return "Hello " + name
}

// Function yang tidak return value
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayHelloTo(from string, to string) {
	fmt.Println("Hello", to+", my name is", from)
}

func helloWord() {
	fmt.Println("Hello World")
}

/*
*
Named return value
Jadi kalo gamau menulis tipe datanya untuk dikembalikan, bisa menggunakan nama variable dan tipe datanya yang akan dikembalikan nanti
Jadi nanti di dalam functionnya tinggal balikin nama variablenya
*/
func getCompleteName() (firstName, lastName string, age int) {
	firstName = "Ahmad"
	lastName = "Raharjo"
	age = 22

	//Balikinnya bisa langsung return gini atau tulis ulang nama2x variablenya
	return
}

/*
*
Variadic function
Ini sama aja kayak di java sih, jadi parameter di akhir function bisa dimasukkan lebih dari 1 value dan nanti menjadi array
*/
func sumAll(numbers ...int) (total int) {
	total = 0

	for _, value := range numbers {
		total += value
	}
	return
}
