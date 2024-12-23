package controller

import (
	"github.com/gofiber/fiber/v2"
    "net/http"
	_productManagingService "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/service"
    "github.com/onizukazaza/tar-ecom-api/pkg/custom"
    _productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type productManagingControllerImpl struct {
	productManagingService _productManagingService.ProductManagingService
}

func NewProductManagingControllerImpl(
	productManagingService _productManagingService.ProductManagingService,
) ProductManagingController {
	return &productManagingControllerImpl{productManagingService}
    //return &productManagingControllerImpl{productService: productService}
}


func (c *productManagingControllerImpl) Listing(ctx *fiber.Ctx) error {

    filter := &_productManagingModel.FilterRequest{}


    customRequest := custom.NewCustomFiberRequest(ctx)
    if err := customRequest.Bind(filter); err != nil {
        return custom.CustomError(ctx, fiber.StatusBadRequest, err.Error())
    }


    productModelList, err := c.productManagingService.Listing(filter)
    if err != nil {
        return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusOK).JSON(productModelList)
}



func (c *productManagingControllerImpl) GetProductByID(ctx *fiber.Ctx) error {
    productID := ctx.Params("id") // ดึง id จาก URL

    product, err := c.productManagingService.ViewProductByID(productID)
    if err != nil {
        // return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
        return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
    }


    return ctx.Status(fiber.StatusOK).JSON(product)
}