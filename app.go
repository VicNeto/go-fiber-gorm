package main

import (
	"go-fiber-template/models"
	"go-fiber-template/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	models.InitDB()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${pid} - ${time} ${locals:requestid} ${status} - ${method} ${path}​\n​",
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})
	app.Get("/:greet", func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("greet"))
	})

	routers.SetupRoutes(app)

	app.Listen(":3007")
}
