package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
    SetRole(ctx *fiber.Ctx) error
	
}