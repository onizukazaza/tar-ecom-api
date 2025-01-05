package controller

import (
	 "github.com/gofiber/fiber/v2"
)

type OAuth2Controller interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error

	UserAuthorizing(ctx *fiber.Ctx) error
	
}

