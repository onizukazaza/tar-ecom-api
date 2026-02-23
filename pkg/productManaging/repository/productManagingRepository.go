package repository

import (
	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type ProductManagingRepository interface {
	Listing(filter *_productManagingModel.FilterRequestBySeller, sellerID string) ([]*_productManagingModel.ProductDetail, error)
	GetProductByID(productID uuid.UUID) (*_productManagingModel.ProductDetail, error) // เพิ่มฟังก์ชันใหม่

	ListActiveProducts(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error)
	GetProductByIDAndSeller(productID uuid.UUID, sellerID string) (*_productManagingModel.ProductDetail, error)

}
