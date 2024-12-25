package controller

import (
	"github.com/gofiber/fiber/v2"
	_productService "github.com/onizukazaza/tar-ecom-api/pkg/product/service"
	_productModel "github.com/onizukazaza/tar-ecom-api/pkg/product/model"
	"net/http"
	"github.com/onizukazaza/tar-ecom-api/pkg/custom"
)

type productControllerImpl struct {
	productService _productService.ProductService
}

func NewProductController(productService _productService.ProductService) ProductController {
	return &productControllerImpl{productService: productService}
}

// CreateProduct handles the creation of a new product
func (c *productControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
    var req _productModel.ProductCreatingReq

    // ใช้ custom request binding
    customRequest := custom.NewCustomFiberRequest(ctx)
    if err := customRequest.Bind(&req); err != nil {
        return custom.CustomError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
    }

    err := c.productService.CreateProduct(&req)
    if err != nil {
        return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}


func (c *productControllerImpl) EditProduct(ctx *fiber.Ctx) error {
	var req _productModel.ProductEditingReq 

	// ดึง Path Parameter "id" และตั้งค่าใน Struct
	if id := ctx.Params("id"); id != "" {
		req.ID = id
	}

	// Bind Request
	customRequest := custom.NewCustomFiberRequest(ctx)
	if err := customRequest.Bind(&req); err != nil {
		return custom.CustomError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	// Call Service
	err := c.productService.EditProduct(&req)
	if err != nil {
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Product updated successfully"})
}



func (c *productControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {

	productID := ctx.Params("id")
	if productID == "" {
		return custom.CustomError(ctx, fiber.StatusBadRequest, "Product ID is required")
	}


	err := c.productService.DeleteProduct(productID)
	if err != nil {
		return custom.CustomError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}
