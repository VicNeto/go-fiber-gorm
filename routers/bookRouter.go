package routers

import (
	"go-fiber-template/controllers"

	"github.com/gofiber/fiber/v2"
)

// AddBookRoutes router
func AddBookRoutes(api fiber.Router) {
	api.Get("/books", controllers.ListBooks)
	api.Post("/books", controllers.NewBook)
	api.Get("/books/:id", controllers.GetBook)
	api.Delete("/books/:id", controllers.DeleteBook)
}
