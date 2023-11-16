package main

import "fmt"

/*
Di bahasa pemorgraman lain jika object belum diberikan data biasanya akan diberikan null
Di go untuk tipe data tertentu data awalnya adalah default valuenya
namun di go juga ada data kosong nil yang bisa digunakan untuk tipe data tertentu seperti
interface, function, map, slice, pointer, dan channel
*/

//Ini tidak bisa karena string kalo kosong isi defaultnya adalah string kosong
//func Contoh(name string) string {
//	if name == "" {
//		return nil
//	}
//}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Datanya kosong")
	} else {
		fmt.Println(data["name"])
	}

	data = NewMap("Gayuh")
	if data == nil {
		fmt.Println("Datanya kosong")
	} else {
		fmt.Println(data["name"])
	}
}
