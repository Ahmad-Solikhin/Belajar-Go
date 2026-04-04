package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Hobbies    []string `json:"hobbies"`
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
