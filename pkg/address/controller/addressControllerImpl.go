package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/onizukazaza/tar-ecom-api/pkg/custom"
    _addressService "github.com/onizukazaza/tar-ecom-api/pkg/address/service"
    _addressModel "github.com/onizukazaza/tar-ecom-api/pkg/address/model"
	"github.com/onizukazaza/tar-ecom-api/pkg/validation"
	
)

type addressControllerImpl struct {
    addressService _addressService.AddressService
}

func NewAddressControllerImpl(service _addressService.AddressService) AddressController {
    return &addressControllerImpl{addressService: service}
}

func (c *addressControllerImpl) CreateAddress(ctx *fiber.Ctx) error {

    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, err.Error())
    }

    req := new(_addressModel.CreateAddressReq)
    customReq := custom.NewCustomFiberRequest(ctx)

    if err := customReq.Bind(req); err != nil {
        return custom.CustomError(ctx, fiber.StatusBadRequest, err.Error())
    }

    req.UserID = buyerID

    createdAddress, err := c.addressService.CreateAddress(req)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusCreated).JSON(createdAddress)
}

func (c *addressControllerImpl) ListAddresses(ctx *fiber.Ctx) error {
    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, "Unauthorized access")
    }

    addresses, err := c.addressService.ListAddresses(buyerID)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusNotFound, err.Error())
    }

    return ctx.Status(fiber.StatusOK).JSON(addresses)
}


func (c *addressControllerImpl) FindAddressByID(ctx *fiber.Ctx) error {
    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, err.Error())
    }

    id := ctx.Params("id")

    address, err := c.addressService.FindAddressByID(id, buyerID)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusNotFound, "Address not found")
    }

    return ctx.Status(fiber.StatusOK).JSON(address)
}

func (c *addressControllerImpl) UpdateFavoriteAddress(ctx *fiber.Ctx) error {
    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, err.Error())
    }

    id := ctx.Params("id")

    var body struct {
        Favorite bool `json:"favorite"`
    }

    if err := ctx.BodyParser(&body); err != nil {
        return custom.CustomError(ctx, fiber.StatusBadRequest, "Invalid request body")
    }

    if err := c.addressService.UpdateFavoriteAddress(id, buyerID, body.Favorite); err != nil {
        return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Favorite status updated successfully",
    })
}

func (c *addressControllerImpl) DeleteAddress(ctx *fiber.Ctx) error {
    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, "Unauthorized access")
    }

    id := ctx.Params("id")

    if err := c.addressService.DeleteAddress(id, buyerID); err != nil {
        return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Address deleted successfully",
    })
}

func (c *addressControllerImpl) EditAddress(ctx *fiber.Ctx) error {
    // ดึง User ID จาก Context (เช่น จาก JWT Token)
    buyerID, err := validation.BuyerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusUnauthorized, err.Error())
    }

    // สร้างโครงสร้างสำหรับรับ Request
    var req _addressModel.EditAddressReq

    // ดึง ID จาก Path Parameter (เช่น /v1/user/address/:id)
    if id := ctx.Params("id"); id != "" {
        req.ID = id
    }

    // ใช้ Custom Request Binder เพื่อตรวจสอบและ Bind ค่า
    customRequest := custom.NewCustomFiberRequest(ctx)
    if err := customRequest.Bind(&req); err != nil {
        return custom.CustomError(ctx, fiber.StatusBadRequest, "Invalid request: "+err.Error())
    }

    // ตั้งค่า UserID ที่ผูกกับที่อยู่
    req.UserID = buyerID

    // เรียก Service Layer เพื่อทำงาน
    err = c.addressService.EditAddress(&req)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
    }

    // ส่ง Response กลับ
    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Address updated successfully"})
}
