package controller

import (
	"github.com/onizukazaza/tar-ecom-api/pkg/validation"
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


func (c *productControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	sellerID, err := validation.SellerIDGetting(ctx)
	if err != nil {

		return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
	}

	var req _productModel.ProductCreatingReq
	customRequest := custom.NewCustomFiberRequest(ctx)
	if err := customRequest.Bind(&req); err != nil {
		return custom.CustomError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	req.SellerID = sellerID
	err = c.productService.CreateProduct(&req)
	if err != nil {
		return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}


func (c *productControllerImpl) EditProduct(ctx *fiber.Ctx) error {
    sellerID, err := validation.SellerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
    }

    var req _productModel.ProductEditingReq
    if id := ctx.Params("id"); id != "" {
        req.ID = id
    }

    customRequest := custom.NewCustomFiberRequest(ctx)
    if err := customRequest.Bind(&req); err != nil {
        return custom.CustomError(ctx, http.StatusBadRequest, "Invalid request: "+err.Error())
    }

    req.SellerID = sellerID
    err = c.productService.EditProduct(&req)
    if err != nil {
        return custom.CustomError(ctx, http.StatusForbidden, err.Error())
    }

    return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Product updated successfully"})
}



func (c *productControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {
    sellerID, err := validation.SellerIDGetting(ctx)
    if err != nil {
        return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
    }

    productID := ctx.Params("id")
    if productID == "" {
        return custom.CustomError(ctx, fiber.StatusBadRequest, "Product ID is required")
    }

    err = c.productService.DeleteProductWithSeller(productID, sellerID)
    if err != nil {
        return custom.CustomError(ctx, http.StatusInternalServerError, err.Error())
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}
