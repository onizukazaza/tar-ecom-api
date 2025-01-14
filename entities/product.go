package entities

import (
	"time"

	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	
)


type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

type Product struct {
	ID               uuid.UUID          `db:"id"`
	ProductName      string             `db:"product_name"`
	Description      string             `db:"description"`
	SellerID         uuid.UUID          `db:"seller_id"`
	Gender           Gender             `db:"gender"`
	IsArchive        bool               `db:"is_archive"`
	CreatedAt        time.Time          `db:"created_at"`
	UpdatedAt        time.Time          `db:"updated_at"`
	ProductImages    *[]ProductImage    `db:"-"`
	ProductVariation *[]ProductVariation `db:"-"`
}

func (p *Product) ToModel(primaryImage _productManagingModel.ImageInfo, additionalImages []_productManagingModel.ImageInfo, variations []_productManagingModel.ProductVariationInfo) *_productManagingModel.ProductDetail {
	return &_productManagingModel.ProductDetail{
		ID:               p.ID.String(),
		ProductName:      p.ProductName,
		Description:      p.Description,
		SellerID:         p.SellerID.String(),
		Gender:           string(p.Gender), 
		IsArchive:        p.IsArchive, 
		PrimaryImage:     primaryImage,
		AdditionalImages: additionalImages,
		Variations:       variations,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

