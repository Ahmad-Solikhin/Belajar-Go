package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayJson(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Hobbies:    []string{"Hobby1", "Hobby2"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
