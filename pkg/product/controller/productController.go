package controller

import "github.com/gofiber/fiber/v2"

// ProductController defines the contract for product-related HTTP handlers
type ProductController interface {
    CreateProduct(ctx *fiber.Ctx) error
    EditProduct(ctx *fiber.Ctx) error
    DeleteProduct(ctx *fiber.Ctx) error
}
