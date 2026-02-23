package entities

import (
	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	"time"
)

type Size struct {
	ID        uuid.UUID `db:"id"`
	SizeType  string    `db:"size_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}


func (s *Size) ToModel() _productManagingModel.SizeInfo {
	return _productManagingModel.SizeInfo{
		ID:       s.ID.String(),
		SizeType: s.SizeType,
	}
}
