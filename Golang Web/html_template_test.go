package Golang_Web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	parse, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "SIMPLE", "Haiyaa Looo")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLGlob(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "simple2.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLGlob(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLGlob(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

//go:embed templates/*.gohtml
var templates embed.FS

func SimpleHTMLFS(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "simple2.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLDataMap(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Hello HTML Template",
		"Name":  "Gayuh",
	})
}

func TestSimpleHTMLDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

type Page struct {
	Title   string
	Name    string
	Address map[string]string
}

func SimpleHTMLDataStruct(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "name.gohtml", Page{
		Name:    "Ahmad",
		Title:   "Hello HTML Template",
		Address: map[string]string{"Country": "Indonesia"},
	})
}

func TestSimpleHTMLDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLIf(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "if.gohtml", Page{
		Name: "",
	})
}

func TestSimpleHTMLIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLComparator(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 50,
	})
}

func TestSimpleHTMLComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLComparator(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLRange(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{"Makan", "Minum"},
	})
}

func TestSimpleHTMLRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLWith(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Address": map[string]string{"Country": "Indonesia", "Street": "Jalan"},
	})
}

func TestSimpleHTMLWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func SimpleHTMLLayout(writer http.ResponseWriter, request *http.Request) {
	parse, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	parse.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Name":  "Raharjo",
		"Title": "Test Layout",
	})
}

func TestSimpleHTMLLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
