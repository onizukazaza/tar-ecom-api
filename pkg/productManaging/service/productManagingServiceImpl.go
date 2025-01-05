package service

import (
	"fmt"

	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
)

type productManagingServiceImpl struct {
	productManagingRepository _productManagingRepository.ProductManagingRepository
}

func NewProductManagingServiceImpl(productManagingRepository _productManagingRepository.ProductManagingRepository) ProductManagingService {
	return &productManagingServiceImpl{
		productManagingRepository: productManagingRepository,
	}
}


func (s *productManagingServiceImpl) Listing(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error) {

    products, err := s.productManagingRepository.Listing(filter)
    if err != nil {
        return nil, err
    }

    return products, nil
}

func (s *productManagingServiceImpl) ViewProductByID(productID string) (*_productManagingModel.ProductDetail, error) {
	uuidProductID, err := uuid.Parse(productID)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID format: %w", err)
	}

	product, err := s.productManagingRepository.GetProductByID(uuidProductID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
