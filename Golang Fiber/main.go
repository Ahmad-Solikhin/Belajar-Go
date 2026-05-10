package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Prefork:      true,
	}

	app := fiber.New(config)

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Middleware before processing request")
		err := c.Next()
		fmt.Println("Middleware after processing request")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return ctx.Send([]byte(`{"name": "Gayuh"}`))
	})
	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return ctx.Send([]byte(`{"name": "Gayuh"}`))
	})

	if fiber.IsChild() {
		fmt.Println("Child process ")
	} else {
		fmt.Println("Parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
