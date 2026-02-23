package server

import (
	_adminRepository "github.com/onizukazaza/tar-ecom-api/pkg/admin/repository"
	_adminService "github.com/onizukazaza/tar-ecom-api/pkg/admin/service"
	_adminController "github.com/onizukazaza/tar-ecom-api/pkg/admin/controller"
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
)

func (s *fiberServer) initAdminRouter(authorizingMiddleware *authorizingMiddleware) {

	router := s.app.Group("/v1/admin", ErrorHandlerMiddleware() , authorizingMiddleware.MiddlewareFunc())

	// Dependency Injection
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db)
	adminService := _adminService.NewAdminServiceImpl(
		adminRepository,
		userRepository,
	)
	adminController := _adminController.NewAdminControllerImpl(adminService)

	//  Endpoint
	router.Post("/user/:id/set-role", adminController.SetRole)

}