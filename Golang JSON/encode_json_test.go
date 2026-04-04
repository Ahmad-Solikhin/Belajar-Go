package Golang_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
	return string(bytes)
}

func TestEncode(t *testing.T) {
	logJson("Gayuh")
	logJson(1)
	logJson(true)
	logJson([]string{"Ahmad", "Solikhin", "Gayuh"})

}
