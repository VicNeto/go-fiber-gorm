package controllers

import "github.com/gofiber/fiber/v2"

// ListUsers controller
func ListUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "users",
	})
}

// GetUser controller
func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "users",
		"id":   c.Params("id"),
	})
}
