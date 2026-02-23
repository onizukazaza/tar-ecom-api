package controller

import (
	 "github.com/gofiber/fiber/v2"
)

type OAuth2Controller interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error

	//Authenticate
	UserAuthorizing(ctx *fiber.Ctx) error
	
}

