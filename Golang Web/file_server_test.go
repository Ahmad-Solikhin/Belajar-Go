package Golang_Web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

/*
Fitur ini bisa membuat handler di go yang digunakan sebagai static file server, jadi tidak perlu manual me-load file lagi
Karena FileServer adalah handler, jadi bisa ditambahkan ke dalam http.Server atau http.ServeMux

Ketika dicoba maka akan dapat error 404 (Not Found), karena file server akan membaca sebagai berikut /resources/static/index.js
Untuk mengatasi hal ini bisa menggunakan function http.StripPrefix
*/

func TestFileServerFail(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static", fileServer)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestFileServerSuccess(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Untuk golang embed akan menambahkan nama rresourrce menjadi url nya, jadinya static/resources/index.html
Untuk mengatasinya harrus menggunakan fs.Sub
*/

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
