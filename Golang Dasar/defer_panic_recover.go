package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
	//Ini recover yang bener
	message := recover()
	fmt.Println(message)
}

func runApp(error bool) {
	defer logging()
	if error {
		panic("Error")
	}
	fmt.Println("Running Application")
}

func wrongRecover(error bool) {
	defer logging()
	if error {
		panic("Error")
	}

	//Ini naro recover yang salah, dan tidak pernah akan dieksekusi jika error
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func main() {
	/**
	defer : ini function yang selalu dieksekusi setelah menjalankan function baik error atau tidak (kayak finally konsepnya)
	panic : function yang digunakan untuk menghentikan program, saat panic dipanggil propgram akan dihentikan namun defer akan tetap dieksekusi (konsepnya mirip throw checked exception)
	recover : function yang digunakan untuk menangkap data panic, jadi proses program akan tetap berjalan (konsepnya mirip try catch)
	*/

	wrongRecover(true)
	runApp(true)

}
