package repository

import (
    "github.com/onizukazaza/tar-ecom-api/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)
type ProductRepository interface {
	CreateProduct(tx *sqlx.Tx, product *entities.Product, images []entities.ProductImage, variations []entities.ProductVariation) error
	// UpdateProduct(productID uuid.UUID, updates map[string]interface{}) error
	EditProduct(productID uuid.UUID, updates map[string]interface{}, images []entities.ProductImage, variations []entities.ProductVariation) error 

	
	DeleteProduct(tx *sqlx.Tx, productID uuid.UUID) error 
	IsProductOwnedBySeller(productID uuid.UUID, sellerID string) (bool, error)

	GetDB() *sqlx.DB
	
}