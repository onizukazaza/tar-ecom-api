package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	EditProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
	Listing(ctx *fiber.Ctx) error
	FindProductByID(ctx *fiber.Ctx) error
	//owned product all (filter : true , false)

}
