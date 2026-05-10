package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engineMustache = mustache.New("./template", ".mustache")

var app = fiber.New(fiber.Config{
	Views: engineMustache,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error " + err.Error())
	},
})

func TestRoutingHelloWorld(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", string(bytes))
}

func TestCtxQueryParameter(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello " + name + "!")
	})

	request := httptest.NewRequest("GET", "/hello?name=Gayuh", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Gayuh!", string(bytes))

	request = httptest.NewRequest("GET", "/hello", nil)
	response, err = app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest!", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Ahmad")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Raharjo"})
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Ahmad Raharjo", string(bytes))
}

func TestRouteParam(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("User : " + userId + ", " + orderId)
	})

	request := httptest.NewRequest("GET", "/users/1/orders/2", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "User : 1, 2", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name + "!")
	})

	body := strings.NewReader("name=Gayuh")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Gayuh!", string(bytes))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := writer.CreateFormFile("file", "contoh.txt")
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := &LoginRequest{}
		err := json.Unmarshal(body, &request)
		if err != nil {
			return err
		}

		return c.SendString("Hello " + request.Username + "!")
	})

	body := strings.NewReader(`{"username": "Gayuh", "password": "123456"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Gayuh!", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := &RegisterRequest{}
		err := c.BodyParser(request)
		if err != nil {
			return err
		}

		return c.SendString("Hello " + request.Username + "!")
	})

	body := strings.NewReader(`{"username": "Gayuh", "password": "123456"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Gayuh!", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "Gayuh",
			"email":    "ahmadsgr39@gmail.com",
		})
	})

	request := httptest.NewRequest("GET", "/user", nil)
	request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"email":"ahmadsgr39@gmail.com","username":"Gayuh"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/contoh.txt", "contoh.txt")
	})

	request := httptest.NewRequest("GET", "/download", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, "attachment; filename=\"contoh.txt\"", response.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Ini adalah file contoh untuk dikirimkan ke target", string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")
	web.Get("/world", helloWorld)
	web.Get("/hello", helloWorld)

	request := httptest.NewRequest("GET", "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	request := httptest.NewRequest("GET", "/public/contoh.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Ini adalah file contoh untuk dikirimkan ke target", string(bytes))
}

func TestError(t *testing.T) {
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("Haiyaa salah loo")
	})

	request := httptest.NewRequest("GET", "/error", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error Haiyaa salah loo", string(bytes))
}

func TestView(t *testing.T) {
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello Mustache",
			"header":  "Ini Header",
			"content": "Ini Content",
		})
	})

	request := httptest.NewRequest("GET", "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	fmt.Println(string(bytes))
	assert.Contains(t, string(bytes), "Hello Mustache")
}

func TestClient(t *testing.T) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	agent := client.Get("https://www.google.com")
	status, response, err := agent.String()
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, status)
	fmt.Println(string(response))

}
