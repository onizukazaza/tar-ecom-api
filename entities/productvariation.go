package entities

import (
	"time"

	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
)

type ProductVariation struct {
	ID              uuid.UUID `db:"id"`
	ProductID       uuid.UUID `db:"product_id"`
	ColorID         uuid.UUID `db:"color_id"`
	SizeID          uuid.UUID `db:"size_id"`
	Color           Color     `db:"-"`
	Size            Size      `db:"-"`
	VariationPrice  float64   `db:"variation_price"`
	Quantity        int       `db:"quantity"`
	ImageVariations string    `db:"image_variation"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`

	Images []string `db:"-"`
}

func (pv *ProductVariation) ToModel(colorType string, sizeType string) _productManagingModel.ProductVariationInfo {
	return _productManagingModel.ProductVariationInfo{
		ID:              pv.ID.String(),
		ColorType:       colorType,
		SizeType:        sizeType,
		VariationPrice:  pv.VariationPrice,
		Quantity:        pv.Quantity,
		ImageVariations: pv.ImageVariations,
	}
}
