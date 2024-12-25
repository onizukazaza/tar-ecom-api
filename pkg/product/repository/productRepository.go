package repository

import (
    "github.com/onizukazaza/tar-ecom-api/entities"
	"github.com/google/uuid"
)
type ProductRepository interface {
	CreateProduct(product *entities.Product, images []entities.ProductImage, variations []entities.ProductVariation) error
	UpdateProduct(productID uuid.UUID, updates map[string]interface{}) error
	

	UpdateProductImages(productID uuid.UUID, images []entities.ProductImage) error
	UpdateProductVariations(productID uuid.UUID, variations []entities.ProductVariation) error
	
	DeleteProduct(productID uuid.UUID) error
}