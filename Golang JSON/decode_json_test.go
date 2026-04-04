package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"first_name":"John","middle_name":"Doe","last_name":"Smith","hobbies":null}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	json.Unmarshal(jsonBytes, &customer)

	fmt.Println(customer)
}
