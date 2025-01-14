package service

import (
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type ProductManagingService interface {
	GetProductByID(productID string) (*_productManagingModel.ProductDetail, error) 
	ListActiveProducts(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error)
	
}
