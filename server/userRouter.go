package server

import (
	// "github.com/gofiber/fiber/v2/middleware/logger"

	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	_userService "github.com/onizukazaza/tar-ecom-api/pkg/user/service"
	_userController "github.com/onizukazaza/tar-ecom-api/pkg/user/controller"
)

func (s *fiberServer) initUserRouter() {
    // s.app.Use(logger.New())
	router := s.app.Group("/v1/user")
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserControllerImpl(userService)

	router.Get("", userController.Listing)
	router.Get("/:id", userController.FindUserByID)
	router.Post("", userController.CreateUser)
	router.Patch("/:id", userController.EditUser)

}
