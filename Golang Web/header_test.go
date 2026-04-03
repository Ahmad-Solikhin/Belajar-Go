package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("content-type")
	w.Header().Set("Content-Type", "Haiyaaa")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=Gayuh&name=Ahmad", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	responseHeader := response.Header.Get("Content-Type")
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
	fmt.Println(responseHeader)
}
