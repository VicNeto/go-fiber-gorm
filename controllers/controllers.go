package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse json error
func ErrorResponse(c *fiber.Ctx, code int, msg string) {
	JSONResponse(c, code, map[string]string{"error": msg})
}

// JSONResponse json
func JSONResponse(c *fiber.Ctx, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	c.JSON(response)
}
