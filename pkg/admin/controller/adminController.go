package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
    Listing(ctx *fiber.Ctx) error
}