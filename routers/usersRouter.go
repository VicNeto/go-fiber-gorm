package routers

import (
	"go-fiber-template/controllers"

	"github.com/gofiber/fiber/v2"
)

// AddUserRoutes router
func AddUserRoutes(api fiber.Router) {
	api.Get("/users", controllers.ListUsers)
	api.Get("/users/:id", controllers.GetUser)
}
