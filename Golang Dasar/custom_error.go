package main

import (
	"fmt"
	"strings"
)

/*
Karena error adalah interface jadi jika ingin membuat error sendiri bisa membuat error dengan implementasi method Error
*/

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func main() {

	err := SaveData("ahaha", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error", validationErr.Error())

		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error", notFoundErr.Error())

		} else {
			fmt.Println("Unknown error", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}

	//Menggunakan switch case
	switch finalError := err.(type) {
	case *validationError:
		fmt.Println("Validation error", finalError.Error())
	case *notFoundError:
		fmt.Println("Not found error", finalError.Error())
	default:
		fmt.Println("Unknown error")

	}

}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Id cannot be empty"}
	}

	if !strings.EqualFold(id, "gayuh") {
		return &notFoundError{"Data not found"}
	}

	return nil
}
