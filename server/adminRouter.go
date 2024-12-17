package server

import (
	"log"
	"os"

	_adminRepository "github.com/onizukazaza/tar-ecom-api/pkg/admin/repository"
	_adminService "github.com/onizukazaza/tar-ecom-api/pkg/admin/service"
	_adminController "github.com/onizukazaza/tar-ecom-api/pkg/admin/controller"
)

func (s *fiberServer) initUserRouter() {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	router := s.app.Group("/v1/admin")
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, logger)
	adminService := _adminService.NewAdminServiceImpl(adminRepository)
	adminController := _adminController.NewAdminControllerImpl(adminService)
	router.Get("", adminController.Listing)
}
