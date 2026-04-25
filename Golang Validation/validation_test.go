package Golang_Validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("validate should not be nil")
	}
}

func TestValidationField(t *testing.T) {
	validate := validator.New()
	user := ""

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia123"

	err := validate.VarWithValue(confirmPassword, password, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "Gayuh12345"

	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "99"

	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

type LoginRequest struct {
	Username string `validate:"required,email"`
	Name     string `validate:"required,min=5"`
}

func TestValidateStruct(t *testing.T) {
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "ahmadsgr39@gmail.com",
		Name:     "Gayuh12345",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "salah",
		Name:     "asgr",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on error tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type RegisterUser struct {
	Username        string `validate:"required,email"`
	Password        string `validate:"required,min=5"`
	ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidateCrossFieldStruct(t *testing.T) {
	validate := validator.New()
	registerUser := RegisterUser{
		Username:        "ahmadsgr39@gmail.com",
		Password:        "Rahasia123",
		ConfirmPassword: "Rahasia12",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type Address struct {
	City    string `validate:"required"`
	Country string `validate:"required"`
}

func TestValidateNestedStruct(t *testing.T) {
	type User struct {
		Name string `validate:"required,min=5"`
		Address
	}

	validate := validator.New()

	address := Address{
		City:    "Bekasi",
		Country: "Indonesia",
	}

	user := User{
		Name:    "Gayuh",
		Address: address,
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateCollectionStruct(t *testing.T) {
	type User struct {
		Name      string    `validate:"required,min=5"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()

	address1 := Address{
		City:    "Bekasi",
		Country: "Indonesia",
	}

	address2 := Address{
		City:    "Jakarta",
		Country: "",
	}

	user := User{
		Name:      "Gayuh",
		Addresses: []Address{address1, address2},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateBasicCollectionStruct(t *testing.T) {
	type User struct {
		Name      string    `validate:"required,min=5"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()

	address1 := Address{
		City:    "Bekasi",
		Country: "Indonesia",
	}

	address2 := Address{
		City:    "Jakarta",
		Country: "",
	}

	user := User{
		Name:      "Gayuh",
		Addresses: []Address{address1, address2},
		Hobbies:   []string{"Sleep", ""},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}
