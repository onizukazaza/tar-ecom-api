package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onizukazaza/tar-ecom-api/entities"
	_productManagingException "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/exception"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type productManagingRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductManagingRepositoryImpl(db *sqlx.DB) *productManagingRepositoryImpl {
	return &productManagingRepositoryImpl{db: db}
}


func (r *productManagingRepositoryImpl) buildProductBaseQuery() string {
	return `
	FROM products p
	LEFT JOIN product_image pi ON pi.product_id = p.id
	LEFT JOIN product_variation pv ON pv.product_id = p.id
	LEFT JOIN color c ON pv.color_id = c.id
	LEFT JOIN size s ON pv.size_id = s.id
	`
}


func (r *productManagingRepositoryImpl) buildProductSelectFields() string {
	return `
	p.id AS product_id, p.product_name, p.description, p.seller_id, p.gender,
	p.is_archive, p.created_at, p.updated_at,
	pi.id AS image_id, pi.image_url, pi.is_primary,
	pv.id AS variation_id, pv.color_id, pv.size_id, pv.variation_price, pv.quantity, pv.image_variation,
	c.id AS color_id, c.color_type,
	s.id AS size_id, s.size_type
	`
}

func (r *productManagingRepositoryImpl) scanProductRow(rows *sqlx.Rows) (*entities.Product, *entities.ProductImage, *entities.ProductVariation, *entities.Color, *entities.Size, error) {
	var (
		product   entities.Product
		image     entities.ProductImage
		variation entities.ProductVariation
		color     entities.Color
		size      entities.Size
	)

	err := rows.Scan(
		&product.ID,
		&product.ProductName,
		&product.Description,
		&product.SellerID,
		&product.Gender,
		&product.IsArchive,
		&product.CreatedAt,
		&product.UpdatedAt,
		&image.ID,
		&image.ImageURL,
		&image.IsPrimary,
		&variation.ID,
		&variation.ColorID,
		&variation.SizeID,
		&variation.VariationPrice,
		&variation.Quantity,
		&variation.ImageVariations,
		&color.ID,
		&color.ColorType,
		&size.ID,
		&size.SizeType,
	)

	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("error scanning row: %w", err)
	}
	return &product, &image, &variation, &color, &size, nil
}

func (r *productManagingRepositoryImpl) GetProductByID(productID uuid.UUID) (*_productManagingModel.ProductDetail, error) {
	query := `
	SELECT ` + r.buildProductSelectFields() + ` ` + r.buildProductBaseQuery() + `
	WHERE p.id = $1;
	`

	rows, err := r.db.Queryx(query, productID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	return r.processRows(rows)
}

func (r *productManagingRepositoryImpl) GetProductByIDAndSeller(productID uuid.UUID, sellerID string) (*_productManagingModel.ProductDetail, error) {
	query := `
	SELECT ` + r.buildProductSelectFields() + ` ` + r.buildProductBaseQuery() + `
	WHERE p.id = $1 AND p.seller_id = $2;
	`

	rows, err := r.db.Queryx(query, productID, sellerID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	return r.processRows(rows)
}

func (r *productManagingRepositoryImpl) Listing(filter *_productManagingModel.FilterRequestBySeller, sellerID string) ([]*_productManagingModel.ProductDetail, error) {
	query := `
	SELECT ` + r.buildProductSelectFields() + ` ` + r.buildProductBaseQuery() + `
	WHERE p.seller_id = :seller_id
	`

	params := map[string]interface{}{
		"seller_id": sellerID,
	}

	if filter.Gender != "" {
		query += " AND p.gender = :gender"
		params["gender"] = filter.Gender
	}

	if filter.IsArchive != nil {
		query += " AND p.is_archive = :is_archive"
		params["is_archive"] = *filter.IsArchive
	}

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	return r.processRowsForList(rows)
}

func (r *productManagingRepositoryImpl) ListActiveProducts(filter *_productManagingModel.FilterRequest) ([]*_productManagingModel.ProductDetail, error) {
	query := `
	SELECT ` + r.buildProductSelectFields() + ` ` + r.buildProductBaseQuery() + `
	WHERE p.is_archive = false
	`

	if filter.Gender != "" {
		query += " AND p.gender = :gender"
	}

	rows, err := r.db.NamedQuery(query, filter)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	return r.processRowsForList(rows)
}

func (r *productManagingRepositoryImpl) processRows(rows *sqlx.Rows) (*_productManagingModel.ProductDetail, error) {
	productMap := make(map[uuid.UUID]*_productManagingModel.ProductDetail)
	imageMap := make(map[string]struct{})
	variationMap := make(map[string]struct{})

	for rows.Next() {
		product, image, variation, color, size, err := r.scanProductRow(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		r.processProductMaps(
			productMap, 
			imageMap, 
			variationMap, 
			product, 
			image, 
			variation, 
			color, 
			size,
		)
	}

	if len(productMap) == 1 {
		for _, productDetail := range productMap {
			return productDetail, nil
		}
	}

	return nil, &_productManagingException.ProductNotFound{}
}

func (r *productManagingRepositoryImpl) processRowsForList(rows *sqlx.Rows) ([]*_productManagingModel.ProductDetail, error) {
	productMap := make(map[uuid.UUID]*_productManagingModel.ProductDetail)
	imageMap := make(map[string]struct{})
	variationMap := make(map[string]struct{})

	for rows.Next() {
		product, image, variation, color, size, err := r.scanProductRow(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		r.processProductMaps(productMap, 
			imageMap, 
			variationMap, 
			product, 
			image, 
			variation, 
			color, 
			size,
		)
	}

	result := make([]*_productManagingModel.ProductDetail, 0, len(productMap))
	for _, product := range productMap {
		result = append(result, product)
	}

	return result, nil
}

func (r *productManagingRepositoryImpl) processProductMaps(
	productMap map[uuid.UUID]*_productManagingModel.ProductDetail,
	imageMap map[string]struct{},
	variationMap map[string]struct{},
	product *entities.Product,
	image *entities.ProductImage,
	variation *entities.ProductVariation,
	color *entities.Color,
	size *entities.Size,
) {
	if _, exists := productMap[product.ID]; !exists {
		productMap[product.ID] = product.ToModel(_productManagingModel.ImageInfo{}, nil, nil)
	}

	r.processProductImages(productMap[product.ID], image, imageMap)
	r.processProductVariations(productMap[product.ID], variation, color, size, variationMap)
}

func (r *productManagingRepositoryImpl) processProductImages(
	product *_productManagingModel.ProductDetail,
	image *entities.ProductImage,
	imageMap map[string]struct{},
) {
	if image != nil && image.ID != uuid.Nil {
		imageModel := image.ToModel()
		if imageModel.IsPrimary {
			product.PrimaryImage = imageModel
		} else if _, exists := imageMap[imageModel.URL]; !exists {
			imageMap[imageModel.URL] = struct{}{}
			product.AdditionalImages = append(product.AdditionalImages, imageModel)
		}
	}
}

func (r *productManagingRepositoryImpl) processProductVariations(
	product *_productManagingModel.ProductDetail,
	variation *entities.ProductVariation,
	color *entities.Color,
	size *entities.Size,
	variationMap map[string]struct{},
) {
	if variation != nil && variation.ID != uuid.Nil {
		variationKey := fmt.Sprintf("%s-%s", variation.ID.String(), variation.ImageVariations)
		if _, exists := variationMap[variationKey]; !exists {
			variationMap[variationKey] = struct{}{}
			variationModel := variation.ToModel(color.ColorType, size.SizeType)
			product.Variations = append(product.Variations, variationModel)
		}
	}
}
