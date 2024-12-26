package controller

import (
	 "github.com/gofiber/fiber/v2"
)

type OAuth2Controller interface {
	UserLogin(ctx *fiber.Ctx) error
	AdminLogin(ctx *fiber.Ctx) error
	UserLoginCallback(ctx *fiber.Ctx) error
	AdminLoginCallback(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error

// Middlewares
//admin buyer seller middleware func
}

