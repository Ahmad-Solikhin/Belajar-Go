package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId/*rest", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Product: %s, items %s, rest %s", params.ByName("id"), params.ByName("itemId"), strings.Replace(params.ByName("rest"), "/", "", -1))
	})

	request := httptest.NewRequest(http.MethodGet, "/products/1/items/2/saffsa", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, "Product: 1, items 2, rest saffsa", string(body))
}
