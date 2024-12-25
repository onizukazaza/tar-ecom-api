package server

import (
	_productController "github.com/onizukazaza/tar-ecom-api/pkg/product/controller"
	_productRepository "github.com/onizukazaza/tar-ecom-api/pkg/product/repository"
	_productService "github.com/onizukazaza/tar-ecom-api/pkg/product/service"
	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
)

func (s *fiberServer) initProductRouter() {
	// Initialize Router
	router := s.app.Group("/products")

	productManagingRepository := _productManagingRepository.NewProductManagingRepositoryImpl(s.db)
	productRepository := _productRepository.NewProductRepositoryImpl(s.db)
	productService := _productService.NewProductServiceImpl(
		productRepository,
		productManagingRepository,
	)
	productController := _productController.NewProductController(productService)

	router.Post("", productController.CreateProduct)
	router.Patch("/:id", productController.EditProduct)
	router.Delete("/:id", productController.DeleteProduct)
}
