package controllers

import (
	"go-fiber-template/models"

	"github.com/gofiber/fiber/v2"
)

// ListBooks controller
func ListBooks(c *fiber.Ctx) error {
	return models.GetBooks(c)
}

// GetBook controller
func GetBook(c *fiber.Ctx) error {
	return models.GetBook(c)
}

// NewBook controller
func NewBook(c *fiber.Ctx) error {
	return models.NewBook(c)

}

// DeleteBook controller
func DeleteBook(c *fiber.Ctx) error {
	return models.DeleteBook(c)
}
