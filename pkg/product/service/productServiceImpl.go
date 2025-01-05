package service

import (
	"fmt"
	"time"
	"database/sql"

	"github.com/google/uuid"
	"github.com/onizukazaza/tar-ecom-api/entities"
	_productModel "github.com/onizukazaza/tar-ecom-api/pkg/product/model"
	_productRepository "github.com/onizukazaza/tar-ecom-api/pkg/product/repository"
	_productManagingRepository "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/repository"
)

type productServiceImpl struct {
	productRepository _productRepository.ProductRepository
	productManagingRepository _productManagingRepository.ProductManagingRepository
}

func NewProductServiceImpl(
	productRepository _productRepository.ProductRepository ,
	productManagingRepository _productManagingRepository.ProductManagingRepository,  // <- replace with your own repository
	) ProductService {
	return &productServiceImpl{
		productRepository ,
		productManagingRepository , 
	}
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

func NewProduct(req *_productModel.ProductCreatingReq) *entities.Product {
	return &entities.Product{
		ID:          uuid.New(),
		ProductName: req.ProductName,
		Description: req.Description,
		SellerID:    uuid.MustParse(req.SellerID),
		Gender:      entities.Gender(req.Gender),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsArchive:   false,
	}
}

func NewProductImages(productID uuid.UUID, req *_productModel.ProductCreatingReq) []entities.ProductImage {
	images := []entities.ProductImage{
		{
			ID:        uuid.New(),
			ProductID: productID,
			ImageURL:  ToNullString(req.PrimaryImage.URL),
			IsPrimary: ToNullBool(true),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, img := range req.AdditionalImages {
		images = append(images, entities.ProductImage{
			ID:        uuid.New(),
			ProductID: productID,
			ImageURL:  ToNullString(img.URL),
			IsPrimary: ToNullBool(false),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	return images
}

func NewProductVariations(productID uuid.UUID, variationsReq []_productModel.ProductVariationCreatingReq) []entities.ProductVariation {
	variations := make([]entities.ProductVariation, len(variationsReq))
	for i, v := range variationsReq {
		variations[i] = entities.ProductVariation{
			ID:             uuid.New(),
			ProductID:      productID,
			ColorID:        uuid.MustParse(v.ColorID),
			SizeID:         uuid.MustParse(v.SizeID),
			VariationPrice: v.VariationPrice,
			Quantity:       v.Quantity,
			ImageVariations: v.ImageVariations,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
	}
	return variations
}

func (s *productServiceImpl) CreateProduct(req *_productModel.ProductCreatingReq) error {
	// เริ่มต้น Transaction
	tx, err := s.productRepository.GetDB().Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// เตรียมข้อมูล
	product := NewProduct(req)
	images := NewProductImages(product.ID, req)
	variations := NewProductVariations(product.ID, req.Variations)

	// ใช้ฟังก์ชัน Repository เดิมผ่าน Transaction
	err = s.productRepository.CreateProduct(tx, product, images, variations)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create product: %w", err)
	}

	// Commit Transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *productServiceImpl) EditProduct(req *_productModel.ProductEditingReq) error {
    productID, err := uuid.Parse(req.ID)
    if err != nil {
        return fmt.Errorf("invalid product ID: %w", err)
    }

    // ตรวจสอบสิทธิ์การเป็นเจ้าของ
    isOwned, err := s.productRepository.IsProductOwnedBySeller(productID, req.SellerID)
    if err != nil {
        return fmt.Errorf("error verifying product ownership: %w", err)
    }
    if !isOwned {
        return fmt.Errorf("unauthorized: you do not own this product")
    }

    // เตรียมข้อมูลสำหรับการอัปเดต
    updates := map[string]interface{}{
        "updated_at": time.Now(),
    }
    if req.ProductName != "" {
        updates["product_name"] = req.ProductName
    }
    if req.Description != "" {
        updates["description"] = req.Description
    }
    if req.Gender != "" {
        updates["gender"] = req.Gender
    }

    // เตรียมข้อมูลรูปภาพและ Variation
    images := []entities.ProductImage{}
    for _, img := range req.AdditionalImages {
        images = append(images, entities.ProductImage{
            ID:        uuid.MustParse(img.ID),
            ProductID: productID,
            ImageURL:  sql.NullString{String: img.ImageURL, Valid: img.ImageURL != ""},
            IsPrimary: sql.NullBool{Bool: img.IsPrimary, Valid: true},
            UpdatedAt: time.Now(),
        })
    }

    variations := []entities.ProductVariation{}
    for _, variation := range req.Variations {
        variations = append(variations, entities.ProductVariation{
            ID:              uuid.MustParse(variation.ID),
            ProductID:       productID,
            VariationPrice:  variation.VariationPrice,
            Quantity:        variation.Quantity,
            ImageVariations: variation.ImageVariations,
            UpdatedAt:       time.Now(),
        })
    }

    // เรียกฟังก์ชันใน Repository
    err = s.productRepository.EditProduct(productID, updates, images, variations)
    if err != nil {
        return fmt.Errorf("failed to edit product: %w", err)
    }

    return nil
}

func (s *productServiceImpl) DeleteProduct(productID string) error {
    // ตรวจสอบว่า productID เป็น UUID ที่ถูกต้อง
    id, err := uuid.Parse(productID)
    if err != nil {
        return fmt.Errorf("invalid product ID: %w", err)
    }

    // เริ่มต้น Transaction
    tx, err := s.productRepository.GetDB().Beginx()
    if err != nil {
        return fmt.Errorf("failed to start transaction: %w", err)
    }

    // เรียก Repository เพื่อลบสินค้า
    err = s.productRepository.DeleteProduct(tx, id) // ส่ง tx เป็นอาร์กิวเมนต์แรก
    if err != nil {
        tx.Rollback() // Rollback เมื่อเกิดข้อผิดพลาด
        if err.Error() == "product not found" {
            return fmt.Errorf("product not found")
        }
        return fmt.Errorf("failed to delete product: %w", err)
    }

    // Commit Transaction
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}


func (s *productServiceImpl) DeleteProductWithSeller(productID string, sellerID string) error {
    // ตรวจสอบและแปลง ID
    id, err := uuid.Parse(productID)
    if err != nil {
        return fmt.Errorf("invalid product ID: %w", err)
    }

    // เริ่มต้น Transaction
    tx, err := s.productRepository.GetDB().Beginx()
    if err != nil {
        return fmt.Errorf("failed to start transaction: %w", err)
    }

    // ตรวจสอบสิทธิ์การเป็นเจ้าของ
    isOwned, err := s.productRepository.IsProductOwnedBySeller(id, sellerID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("error verifying product ownership: %w", err)
    }
    if !isOwned {
        tx.Rollback()
        return fmt.Errorf("unauthorized: you do not own this product")
    }

    // ลบสินค้า
    err = s.productRepository.DeleteProduct(tx, id) // เพิ่ม `tx` เป็นอาร์กิวเมนต์แรก
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete product: %w", err)
    }

    // Commit Transaction
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}
