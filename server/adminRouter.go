package server

import (
	"github.com/gofiber/fiber/v2/middleware/logger"

	_adminRepository "github.com/onizukazaza/tar-ecom-api/pkg/admin/repository"
	_adminService "github.com/onizukazaza/tar-ecom-api/pkg/admin/service"
	_adminController "github.com/onizukazaza/tar-ecom-api/pkg/admin/controller"
)

func (s *fiberServer) initUserRouter() {
    s.app.Use(logger.New())
	router := s.app.Group("/v1/admin")
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db)
	adminService := _adminService.NewAdminServiceImpl(adminRepository)
	adminController := _adminController.NewAdminControllerImpl(adminService)
	router.Get("", adminController.Listing)
}
