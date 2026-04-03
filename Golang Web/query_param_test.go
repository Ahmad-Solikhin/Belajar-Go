package Golang_Web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=Gayuh", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleParameterValues(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query()["name"]
	fmt.Fprintf(w, "Hello %s", strings.Join(name, ",  "))
}

func TestQueryParameterMultipleValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=Gayuh&name=Ahmad", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
