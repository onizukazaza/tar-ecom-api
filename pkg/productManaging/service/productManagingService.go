package service

import (
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type ProductManagingService interface {
	// Listing() ([]*_productManagingModel.ProductDetail, error)
	 Listing(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error) 
	ViewProductByID(productID string) (*_productManagingModel.ProductDetail, error) // เพิ่มฟังก์ชันใหม่
}
