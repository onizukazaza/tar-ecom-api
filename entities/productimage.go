package entities

import (
	"database/sql"
	"github.com/google/uuid"
	_productManagingModel "github.com/onizukazaza/tar-ecom-api/pkg/productManaging/model"
	"time"
)

type ProductImage struct {
	ID          uuid.UUID      `db:"id"`
	ProductID   uuid.UUID      `db:"product_id"`
	VariationID *uuid.UUID     `db:"variation_id"`
	ImageURL    sql.NullString `db:"image_url"`
	IsPrimary   sql.NullBool   `db:"is_primary"` // รองรับ NULL
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
func (pi *ProductImage) ToModel() _productManagingModel.ImageInfo {
    return _productManagingModel.ImageInfo{
        URL:       safeString(pi.ImageURL),
        IsPrimary: safeBool(pi.IsPrimary),
    }
}

// Helper functions
func safeString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return ""
}

func safeBool(nb sql.NullBool) bool {
    if nb.Valid {
        return nb.Bool
    }
    return false
}
