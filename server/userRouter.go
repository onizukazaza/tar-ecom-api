package server

import (
	// "github.com/gofiber/fiber/v2/middleware/logger"

	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	_userService "github.com/onizukazaza/tar-ecom-api/pkg/user/service"
	_userController "github.com/onizukazaza/tar-ecom-api/pkg/user/controller"
)

func (s *fiberServer) initUserRouter(authorizingMiddleware *authorizingMiddleware) {

    router := s.app.Group("/v1/user", ErrorHandlerMiddleware())

    userRepository := _userRepository.NewUserRepositoryImpl(s.db)
    userService := _userService.NewUserServiceImpl(userRepository)
    userController := _userController.NewUserControllerImpl(userService)

    
    router.Post("", userController.CreateUser) 
    router.Get("", authorizingMiddleware.MiddlewareFunc(), userController.Listing)    // use global middleware role
    router.Get("/:id", authorizingMiddleware.MiddlewareFunc(), userController.FindUserByID) // use global middleware role 
    router.Patch("/edit", authorizingMiddleware.MiddlewareFunc(), userController.EditUser)  // use global middleware role 
}

