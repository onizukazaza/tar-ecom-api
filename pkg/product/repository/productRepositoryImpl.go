package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onizukazaza/tar-ecom-api/entities"
)

type productRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepositoryImpl(db *sqlx.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) CreateProduct(product *entities.Product, images []entities.ProductImage, variations []entities.ProductVariation) error {

	queryProduct := `
		INSERT INTO products (id, product_name, description, seller_id, gender, created_at, updated_at, is_archive)
		VALUES (:id, :product_name, :description, :seller_id, :gender, :created_at, :updated_at, :is_archive)`
	_, err := r.db.NamedExec(queryProduct, product)
	if err != nil {
		return fmt.Errorf("failed to insert product: %w", err)
	}


	for _, img := range images {
		queryImage := `
			INSERT INTO product_image (id, product_id, image_url, is_primary, created_at, updated_at)
			VALUES (:id, :product_id, :image_url, :is_primary, :created_at, :updated_at)`
		_, err := r.db.NamedExec(queryImage, img)
		if err != nil {
			return fmt.Errorf("failed to insert product image: %w", err)
		}
	}


	for _, variation := range variations {
		queryVariation := `
			INSERT INTO product_variation (id, product_id, color_id, size_id, variation_price, quantity, image_variation, created_at, updated_at)
			VALUES (:id, :product_id, :color_id, :size_id, :variation_price, :quantity, :image_variation, :created_at, :updated_at)`
		_, err := r.db.NamedExec(queryVariation, variation)
		if err != nil {
			return fmt.Errorf("failed to insert product variation: %w", err)
		}
	}

	return nil
}

func (r *productRepositoryImpl) UpdateProduct(productID uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE products SET "
	args := []interface{}{}
	i := 1

	for key, value := range updates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}

	query = query[:len(query)-2] // ลบ ", " ท้ายสุด
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, productID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (r *productRepositoryImpl) UpdateProductImages(productID uuid.UUID, images []entities.ProductImage) error {
	query := `
		UPDATE product_image 
		SET image_url = :image_url, 
		    is_primary = :is_primary, 
		    updated_at = :updated_at 
		WHERE id = :id AND product_id = :product_id
	`

	for _, img := range images {
		img.ProductID = productID // เชื่อมโยง product_id กับ image
		_, err := r.db.NamedExec(query, img)
		if err != nil {
			return fmt.Errorf("failed to update product image ID %s: %w", img.ID.String(), err)
		}
	}

	return nil
}

func (r *productRepositoryImpl) UpdateProductVariations(productID uuid.UUID, variations []entities.ProductVariation) error {
	query := `UPDATE product_variation 
              SET variation_price = :variation_price, 
                  quantity = :quantity, 
                  image_variation = :image_variation, 
                  updated_at = :updated_at 
              WHERE id = :id AND product_id = :product_id`

	for _, variation := range variations {
		_, err := r.db.NamedExec(query, variation)
		if err != nil {
			return fmt.Errorf("failed to update product variation: %w", err)
		}
	}

	return nil
}

func (r *productRepositoryImpl) DeleteProduct(productID uuid.UUID) error {
	// ลบข้อมูลใน product_image
	queryImages := `DELETE FROM product_image WHERE product_id = $1`
	_, err := r.db.Exec(queryImages, productID)
	if err != nil {
		return fmt.Errorf("failed to delete product images: %w", err)
	}

	// ลบข้อมูลใน product_variation
	queryVariations := `DELETE FROM product_variation WHERE product_id = $1`
	_, err = r.db.Exec(queryVariations, productID)
	if err != nil {
		return fmt.Errorf("failed to delete product variations: %w", err)
	}

	// ลบข้อมูลใน products
	queryProduct := `DELETE FROM products WHERE id = $1`
	res, err := r.db.Exec(queryProduct, productID)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	// ตรวจสอบว่ามี product ที่ถูกลบหรือไม่
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
