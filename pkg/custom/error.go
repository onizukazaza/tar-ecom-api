package custom

import (
    "github.com/gofiber/fiber/v2"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func CustomError(ctx *fiber.Ctx  , statusCode int, message string ) error {
	return ctx.Status(statusCode).JSON(&ErrorMessage{Message: message})
}