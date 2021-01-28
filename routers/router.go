package routers

import "github.com/gofiber/fiber/v2"

// SetupRoutes routes
func SetupRoutes(app *fiber.App) *fiber.App {
	api := app.Group("/api")
	AddBookRoutes(api)
	AddUserRoutes(api)
	return app
}
