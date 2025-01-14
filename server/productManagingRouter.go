package server

import (
	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
	_productManagingService "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/service"
	_productManagingController "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/controller"
)

func (s *fiberServer) initProductManagingRouter() {

    router := s.app.Group("/items-product", ErrorHandlerMiddleware())
	
	productManagingRepository := _productManagingRepository.NewProductManagingRepositoryImpl(s.db)
	productManagingService := _productManagingService.NewProductManagingServiceImpl(productManagingRepository)
	productManagingController := _productManagingController.NewProductManagingControllerImpl(productManagingService)

	
	//feature for buyer list
	router.Get("/:id", productManagingController.GetProductByID)
	router.Get("/all/buyer", productManagingController.ListActiveProducts)  // user testing
}
