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
	product := NewProduct(req)
	images := NewProductImages(product.ID, req)
	variations := NewProductVariations(product.ID, req.Variations)

	err := s.productRepository.CreateProduct(product, images, variations)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}



func (s *productServiceImpl) EditProduct(req *_productModel.ProductEditingReq) error {
	productID, err := uuid.Parse(req.ID)
	if err != nil {
		return fmt.Errorf("invalid product ID: %w", err)
	}

	// Fetch existing product
	existingProduct, err := s.productManagingRepository.GetProductByID(productID)
	if err != nil {
		return fmt.Errorf("failed to fetch existing product: %w", err)
	}

	// Update product details
	updates := map[string]interface{}{}
	if req.ProductName != "" && req.ProductName != existingProduct.ProductName {
		updates["product_name"] = req.ProductName
	}
	if req.Description != "" && req.Description != existingProduct.Description {
		updates["description"] = req.Description
	}
	if req.Gender != "" && req.Gender != existingProduct.Gender {
		updates["gender"] = req.Gender
	}
	updates["updated_at"] = time.Now()

	err = s.productRepository.UpdateProduct(productID, updates)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	// Update product images
	if len(req.AdditionalImages) > 0 {
		err := s.updateProductImages(productID, req.AdditionalImages)
		if err != nil {
			return err
		}
	}

	// Update product variations
	if len(req.Variations) > 0 {
		err := s.updateProductVariations(productID, req.Variations)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *productServiceImpl) updateProductImages(productID uuid.UUID, images []_productModel.ProductImageUpdatingReq) error {
	for _, image := range images {
		err := s.productRepository.UpdateProductImages(productID, []entities.ProductImage{
			{
				ID:        uuid.MustParse(image.ID),
				ProductID: productID,
				ImageURL:  sql.NullString{String: image.ImageURL, Valid: image.ImageURL != ""},
				IsPrimary: sql.NullBool{Bool: image.IsPrimary, Valid: true},
				UpdatedAt: time.Now(),
			},
		})
		if err != nil {
			return fmt.Errorf("failed to update product images: %w", err)
		}
	}
	return nil
}

func (s *productServiceImpl) updateProductVariations(productID uuid.UUID, variations []_productModel.ProductVariationUpdatingReq) error {
	for _, variation := range variations {
		// แปลงข้อมูลที่อนุญาตให้แก้ไข
		updateData := entities.ProductVariation{
			ID:              uuid.MustParse(variation.ID), // ต้องมี ID
			ProductID:       productID,
			VariationPrice:  variation.VariationPrice,
			Quantity:        variation.Quantity,
			ImageVariations: variation.ImageVariations,
			UpdatedAt:       time.Now(),
		}

		// ส่งข้อมูลไปยัง Repository เพื่ออัปเดต
		err := s.productRepository.UpdateProductVariations(productID, []entities.ProductVariation{updateData})
		if err != nil {
			return fmt.Errorf("failed to update product variations: %w", err)
		}
	}
	return nil
}

func (s *productServiceImpl) DeleteProduct(productID string) error {
	// ตรวจสอบว่า productID เป็น UUID ที่ถูกต้อง
	id, err := uuid.Parse(productID)
	if err != nil {
		return fmt.Errorf("invalid product ID: %w", err)
	}

	// เรียก Repository เพื่อลบสินค้า
	err = s.productRepository.DeleteProduct(id)
	if err != nil {
		if err.Error() == "product not found" {
			return fmt.Errorf("product not found")
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
