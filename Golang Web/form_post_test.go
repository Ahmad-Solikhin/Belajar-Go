package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPostFirstWay(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	form := r.PostForm
	fmt.Fprint(w, form)
}

func FormPostSecondWay(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("firstName")
	lastName := r.PostFormValue("lastName")

	fmt.Fprintf(w, "My name is: %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Gayuh&lastName=Raharjo")
	request := httptest.NewRequest(http.MethodPost, "/", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPostFirstWay(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

	recorder = httptest.NewRecorder()
	FormPostSecondWay(recorder, request)

	response = recorder.Result()
	body, _ = io.ReadAll(response.Body)
	bodyString = string(body)
	fmt.Println(bodyString)
}
