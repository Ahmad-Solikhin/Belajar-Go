package test

import (
	"database/sql"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() http.Handler {
	db := SetupTestDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func SetupTestDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}

func TestCreateCategorySuccess(t *testing.T) {
	router := SetupRouter()

	requestBody := strings.NewReader(`{"name":"test"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {

}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestGetAllCategory(t *testing.T) {

}

func TestUnauthorize(t *testing.T) {

}
