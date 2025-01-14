package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onizukazaza/tar-ecom-api/entities"
	_ProductException "github.com/onizukazaza/tar-ecom-api/pkg/product/exception"
)

type productRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepositoryImpl(db *sqlx.DB) *productRepositoryImpl {
	return &productRepositoryImpl{
		db: db,
	}
}

func (r *productRepositoryImpl) GetDB() *sqlx.DB {
	return r.db
}

func (r *productRepositoryImpl) CreateProduct(tx *sqlx.Tx, product *entities.Product, images []entities.ProductImage, variations []entities.ProductVariation) error {
	_, err := tx.NamedExec(`
		INSERT INTO products (id, product_name, description, seller_id, gender, created_at, updated_at, is_archive)
		VALUES (:id, :product_name, :description, :seller_id, :gender, :created_at, :updated_at, :is_archive)`, product)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}

	if err := r.insertImages(tx, images); err != nil {
		return err
	}

	if err := r.insertVariations(tx, variations); err != nil {
		return err
	}

	return nil
}

func (r *productRepositoryImpl) insertImages(tx *sqlx.Tx, images []entities.ProductImage) error {
	for _, img := range images {
		_, err := tx.NamedExec(`
			INSERT INTO product_image (id, product_id, image_url, is_primary, created_at, updated_at)
			VALUES (:id, :product_id, :image_url, :is_primary, :created_at, :updated_at)`, img)
		if err != nil {
			return fmt.Errorf("failed to insert product image: %w", err)
		}
	}
	return nil
}

func (r *productRepositoryImpl) insertVariations(tx *sqlx.Tx, variations []entities.ProductVariation) error {
	for _, variation := range variations {
		_, err := tx.NamedExec(`
			INSERT INTO product_variation (id, product_id, color_id, size_id, variation_price, quantity, image_variation, created_at, updated_at)
			VALUES (:id, :product_id, :color_id, :size_id, :variation_price, :quantity, :image_variation, :created_at, :updated_at)`, variation)
		if err != nil {
			return &_ProductException.UnCreateProduct{}
		}
	}
	return nil
}

func (r *productRepositoryImpl) EditProduct(productID uuid.UUID, updates map[string]interface{}, images []entities.ProductImage, variations []entities.ProductVariation) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	if err := r.updateProduct(tx, productID, updates); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.updateImages(tx, images); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.updateVariations(tx, variations); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *productRepositoryImpl) updateProduct(tx *sqlx.Tx, productID uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE products SET "
	args := []interface{}{}
	i := 1
	for key, value := range updates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, productID)

	_, err := tx.Exec(query, args...)
	if err != nil {
		return &_ProductException.FailedToUpdateProduct{}
	}
	return nil
}

func (r *productRepositoryImpl) updateImages(tx *sqlx.Tx, images []entities.ProductImage) error {
	for _, img := range images {
		query := `
			UPDATE product_image 
			SET image_url = :image_url, 
				is_primary = :is_primary, 
				updated_at = :updated_at 
			WHERE id = :id AND product_id = :product_id`
		_, err := tx.NamedExec(query, img)
		if err != nil {
			return &_ProductException.FailedToUpdateProductImage{ImageID: img.ID.String()}
		}
	}
	return nil
}

func (r *productRepositoryImpl) updateVariations(tx *sqlx.Tx, variations []entities.ProductVariation) error {
	for _, variation := range variations {
		query := `
			UPDATE product_variation 
			SET variation_price = :variation_price, 
				quantity = :quantity, 
				image_variation = :image_variation, 
				updated_at = :updated_at 
			WHERE id = :id AND product_id = :product_id`
		_, err := tx.NamedExec(query, variation)
		if err != nil {
			return &_ProductException.FailedToUpdateProductVariation{VariationID: variation.ID.String()}
		}
	}
	return nil
}

func (r *productRepositoryImpl) ArchiveProduct(tx *sqlx.Tx, productID uuid.UUID) error {
	query := `
		UPDATE products
		SET is_archive = true, updated_at = NOW()
		WHERE id = $1
	`
	_, err := tx.Exec(query, productID)
	if err != nil {
		return &_ProductException.UnArchive{}
	}
	return nil
}

func (r *productRepositoryImpl) IsProductOwnedBySeller(productID uuid.UUID, sellerID string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM products
			WHERE id = $1 AND seller_id = $2
		)
	`
	err := r.db.Get(&exists, query, productID, sellerID)
	if err != nil {
		return false, fmt.Errorf("failed to check product ownership: %w", err)
	}
	return exists, nil
}


