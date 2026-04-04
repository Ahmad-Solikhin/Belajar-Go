package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"p001", "name":"Samsung", "price":20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"].(string))
	fmt.Println(result["name"].(string))
	fmt.Println(result["price"].(float64))
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "p001",
		"name":  "Samsung",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
