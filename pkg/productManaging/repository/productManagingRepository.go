package repository

import (
	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type ProductManagingRepository interface {
	Listing(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error)
	GetProductByID(productID uuid.UUID) (*_productManagingModel.ProductDetail, error) // เพิ่มฟังก์ชันใหม่
	
	
}
