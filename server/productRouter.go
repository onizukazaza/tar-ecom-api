package server

import (
	_productController "github.com/onizukazaza/tar-ecom-api/pkg/product/controller"
	_productRepository "github.com/onizukazaza/tar-ecom-api/pkg/product/repository"
	_productService "github.com/onizukazaza/tar-ecom-api/pkg/product/service"
	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
	
)

func (s *fiberServer) initProductRouter(authorizingMiddleware *authorizingMiddleware) {
	router := s.app.Group("/product-managing", ErrorHandlerMiddleware(), authorizingMiddleware.MiddlewareFunc())

	// Dependency Injection
	productManagingRepository := _productManagingRepository.NewProductManagingRepositoryImpl(s.db)
	productRepository := _productRepository.NewProductRepositoryImpl(s.db)
	productService := _productService.NewProductServiceImpl(
		productRepository,
		productManagingRepository,
	)
	productController := _productController.NewProductController(productService)

	
	//  Endpoint get product with owner
	router.Post("", productController.CreateProduct)

	router.Get("/all", productController.Listing)  //for owner product management is feature
	router.Get("/seller-id/:id", productController.FindProductByID) //for owner product management is feature
	router.Patch("/:id", productController.EditProduct) //for owner product management is feature
	router.Delete("/:id", productController.DeleteProduct) //for owner product management is feature
}
