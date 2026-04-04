package Golang_JSON

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("test.json")
	encoder := json.NewEncoder(writer)
	defer writer.Close()

	customer := Customer{
		FirstName: "John",
		LastName:  "Doe",
	}

	encoder.Encode(customer)
}
