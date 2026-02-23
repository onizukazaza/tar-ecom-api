package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onizukazaza/tar-ecom-api/pkg/custom"
	_productModel "github.com/onizukazaza/tar-ecom-api/pkg/product/model"
	_productService "github.com/onizukazaza/tar-ecom-api/pkg/product/service"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	"github.com/onizukazaza/tar-ecom-api/pkg/validation"
	"net/http"
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

func (c *productControllerImpl) Listing(ctx *fiber.Ctx) error {

	sellerID, err := validation.SellerIDGetting(ctx)
	if err != nil {
		return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
	}

	filter := &_productManagingModel.FilterRequestBySeller{}
	customRequest := custom.NewCustomFiberRequest(ctx)

	if err := customRequest.Bind(filter); err != nil {
		if err.Error() != "failed to parse JSON body: invalid JSON format" {
			return custom.CustomError(ctx, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
		}
	}

	
	products, err := c.productService.Listing(filter, sellerID)
	if err != nil {
		return custom.CustomError(ctx, fiber.StatusInternalServerError, "Failed to fetch product list: "+err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (c *productControllerImpl) FindProductByID(ctx *fiber.Ctx) error {
	sellerID, err := validation.SellerIDGetting(ctx)
	if err != nil {
		return custom.CustomError(ctx, http.StatusUnauthorized, err.Error())
	}

	productID := ctx.Params("id")
	if productID == "" {
		return custom.CustomError(ctx, fiber.StatusBadRequest, "Product ID is required")
	}

	product, err := c.productService.GetProductByIDAndSeller(productID, sellerID)
	if err != nil {

		return custom.CustomError(ctx, fiber.StatusNotFound, "Product not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}
