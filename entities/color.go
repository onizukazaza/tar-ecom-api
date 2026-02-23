package entities

import (
	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	"time"
)

type Color struct {
	ID        uuid.UUID `db:"id"`
	ColorType string    `db:"color_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}


func (c *Color) ToModel() _productManagingModel.ColorInfo {
	return _productManagingModel.ColorInfo{
		ID:        c.ID.String(),
		ColorType: c.ColorType,
	}
}