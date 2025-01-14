package service

import "github.com/onizukazaza/tar-ecom-api/pkg/address/model"

type AddressService interface {
    CreateAddress(req *model.CreateAddressReq) (*model.Address, error)
    EditAddress(req *model.EditAddressReq) error
    ListAddresses(userID string) ([]*model.Address, error)
    FindAddressByID(id string, userID string) (*model.Address, error)
    UpdateFavoriteAddress(id string, userID string, favorite bool) error
    DeleteAddress(id string, userID string) error

}

// EditAddress(req *model.EditAddressReq) error

// FindAddressByID(id string) (*model.Address, error)
