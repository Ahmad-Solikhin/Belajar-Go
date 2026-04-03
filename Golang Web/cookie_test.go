package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-SECRET"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-SECRET")
	if err != nil {
		fmt.Fprint(w, "Get cookie fail")
	} else {
		fmt.Fprint(w, cookie)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/?name=Gayuh", nil)
	recoreder := httptest.NewRecorder()

	SetCookie(recoreder, request)

	cookies := recoreder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-SECRET"
	cookie.Value = "Gayuh"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", SetCookie)
	mux.HandleFunc("/get", GetCookie)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
