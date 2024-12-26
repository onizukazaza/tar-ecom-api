package controller

import (
    "github.com/gofiber/fiber/v2"
)
type UserController interface {
	Listing(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	FindUserByID(ctx *fiber.Ctx) error
	EditUser(ctx *fiber.Ctx) error
}