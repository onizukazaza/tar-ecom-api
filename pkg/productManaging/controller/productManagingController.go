package controller

import (
	"github.com/gofiber/fiber/v2"
)
type ProductManagingController interface{
	// Listing(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error 
	ListActiveProducts(ctx *fiber.Ctx) error  //testing
}

