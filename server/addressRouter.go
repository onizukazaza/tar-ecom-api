package server

import (
	_addressRepository "github.com/onizukazaza/tar-ecom-api/pkg/address/repository"
	_addressService "github.com/onizukazaza/tar-ecom-api/pkg/address/service"
	_addressController "github.com/onizukazaza/tar-ecom-api/pkg/address/controller"
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
)

func (s *fiberServer) initAddressRouter(authorizingMiddleware *authorizingMiddleware) {

	router := s.app.Group("/v1/user/address", ErrorHandlerMiddleware() , authorizingMiddleware.MiddlewareFunc())

	// Dependency Injection
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	addressRepository := _addressRepository.NewAddressRepositoryImpl(s.db)
	addressService := _addressService.NewAddressServiceImpl(
		addressRepository,
		userRepository,
	)
	addressController := _addressController.NewAddressControllerImpl(addressService)

	// Endpoint
	router.Post("", addressController.CreateAddress)
	router.Get("/all" , addressController.ListAddresses)
	router.Get("/:id", addressController.FindAddressByID)
	router.Patch("/:id", addressController.EditAddress) 
	router.Patch("/:id/favorite", addressController.UpdateFavoriteAddress)
	router.Delete("/:id", addressController.DeleteAddress)
	

}

