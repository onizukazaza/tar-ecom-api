package repository

import "github.com/onizukazaza/tar-ecom-api/entities"

type AddressRepository interface {
    CreateAddress(address *entities.Address) error
    EditAddress(address *entities.Address) error
    ListAddresses(userID string) ([]*entities.Address, error)
    FindAddressByID(id string, userID string) (*entities.Address, error) 
    UpdateFavoriteAddress(id string, userID string, favorite bool) error 
    ClearAllFavorites(userID string) error
    DeleteAddress(id string, userID string) error

}
