package Golang_Validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func TestValidateMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Name      string            `validate:"required,min=5"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`
	}

	validate := validator.New()

	address1 := Address{
		City:    "Bekasi",
		Country: "Indonesia",
	}

	address2 := Address{
		City:    "Jakarta",
		Country: "Indonesia",
	}

	user := User{
		Name:      "Gayuh",
		Addresses: []Address{address1, address2},
		Hobbies:   []string{"Sleep", "Read"},
		Schools: map[string]School{
			"1sda": {Name: "Gayuh"},
			"h":    {Name: ""},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateBasicMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Name      string            `validate:"required,min=5"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`
		Wallet    map[string]int    `validate:"required,dive,keys,required,endkeys,gt=1000"`
	}

	validate := validator.New()

	address1 := Address{
		City:    "Bekasi",
		Country: "Indonesia",
	}

	address2 := Address{
		City:    "Jakarta",
		Country: "Indonesia",
	}

	user := User{
		Name:      "Gayuh",
		Addresses: []Address{address1, address2},
		Hobbies:   []string{"Sleep", "Read"},
		Schools: map[string]School{
			"1sda": {Name: "Gayuh"},
			"h":    {Name: ""},
		},
		Wallet: map[string]int{
			"kosong": 1000,
		},
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "132",
		Name:   "Gayuh",
		Owner:  "Gayuh",
		Slogan: "Hidup",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidation(t *testing.T) {

	validate := validator.New()
	_ = validate.RegisterValidation("username", MustValidUsername)

	type Seller struct {
		Username string `validate:"username"`
	}

	seller := Seller{
		Username: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile(`^[0-9]+$`)

func MustValidPin(field validator.FieldLevel) bool {
	param, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == param
}

func TestCustomValidationWithParam(t *testing.T) {

	validate := validator.New()
	_ = validate.RegisterValidation("pin", MustValidPin)

	type Seller struct {
		Pin string `validate:"pin=10"`
	}

	seller := Seller{
		Pin: "1234543790",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
	}

	validate := validator.New()

	login1 := Login{
		Username: "1234543790",
	}

	err := validate.Struct(login1)
	if err != nil {
		fmt.Println(err.Error())
	}

	login2 := Login{
		Username: "ahmadsgr39@gmail.com",
	}

	err = validate.Struct(login2)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Filed not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("field_equals_ignbore_case", MustEqualsIgnoreCase)
	if err != nil {
		return
	}

	type User struct {
		Username string `validate:"required,field_equals_ignbore_case=Email|field_equals_ignbore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
	}

	user := User{
		Username: "ahmadsgr39@gmail.com",
		Email:    "ahmadsgr39@gmail.com",
		Phone:    "1234543790",
	}

	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// Sukses
	} else {
		// gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "1234543790",
		Email:    "ahmadsgr39@gmail.com",
		Phone:    "1234543790",
		Password: "1234543790",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
