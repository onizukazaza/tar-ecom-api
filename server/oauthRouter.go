package server

import (
	_oauthService "github.com/onizukazaza/tar-ecom-api/pkg/oauth2/service"
	_oauthController "github.com/onizukazaza/tar-ecom-api/pkg/oauth2/controller"
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	
)

func (s *fiberServer) initAuthRouter(authorizingMiddleware *authorizingMiddleware) {
	
	router := s.app.Group("/v1/auth", ErrorHandlerMiddleware())

	// Dependency Injection
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	oauthService := _oauthService.NewOAuth2Service(
		userRepository, 
		s.secretKey,
	)
	oauthController := _oauthController.NewOAuth2Controller(
		oauthService,
		s.secretKey,
		)

    //  Endpoint
	router.Post("/login", oauthController.Login)
	router.Post("/logout",authorizingMiddleware.MiddlewareFunc(), oauthController.Logout)
}
