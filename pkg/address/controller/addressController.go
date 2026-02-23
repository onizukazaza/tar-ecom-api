package controller

import "github.com/gofiber/fiber/v2"

type AddressController interface {
    CreateAddress(ctx *fiber.Ctx) error
    EditAddress(ctx *fiber.Ctx) error
	ListAddresses(ctx *fiber.Ctx) error
    FindAddressByID(ctx *fiber.Ctx) error
    UpdateFavoriteAddress(ctx *fiber.Ctx) error
    DeleteAddress(ctx *fiber.Ctx) error 
    // EditAddress(ctx *fiber.Ctx) error
}