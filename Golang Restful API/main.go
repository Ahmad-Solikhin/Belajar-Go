package main

import (
	"fmt"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:4000",
		Handler: authMiddleware,
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func main() {
	server := InitializeServer()

	fmt.Println("Running server on " + server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
