package server

import (
	// "github.com/gofiber/fiber/v2/middleware/logger"

	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
	_productManagingService "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/service"
	_productManagingController "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/controller"
)

func (s *fiberServer) initProductManagingRouter() {
    // s.app.Use(logger.New())

	// Initialize Router
	router := s.app.Group("/items-product")

	// Initialize dependencies
	productManagingRepository := _productManagingRepository.NewProductManagingRepositoryImpl(s.db)
	productManagingService := _productManagingService.NewProductManagingServiceImpl(productManagingRepository)
	productManagingController := _productManagingController.NewProductManagingControllerImpl(productManagingService)

	// Define routes
	router.Get("", productManagingController.Listing)        // Route สำหรับ listing สินค้าทั้งหมด
	router.Get("/:id", productManagingController.GetProductByID) // Route สำหรับค้นหาสินค้าตาม ID
}
