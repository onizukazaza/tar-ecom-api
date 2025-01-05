package server

import (
    "github.com/gofiber/fiber/v2"
)

func ErrorHandlerMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        defer func() {
            if r := recover(); r != nil {
                c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "error": "Internal Server Error",
                    "detail": r,
                })
            }
        }()
        return c.Next()
    }
}
