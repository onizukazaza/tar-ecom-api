package model

import "time"

type (
	ProductDetail struct {
		ID               string                 `json:"id"`
		ProductName      string                 `json:"product_name"`
		Description      string                 `json:"description"`
		SellerID         string                 `json:"seller_id"`
		Gender           string                 `json:"gender"`
		Variations       []ProductVariationInfo `json:"variations"`
		PrimaryImage     ImageInfo              `json:"primary_image"`
		AdditionalImages []ImageInfo            `json:"additional_images"`
		IsArchive        bool                   `json:"is_archive"`
		CreatedAt        time.Time              `json:"created_at"`
		UpdatedAt        time.Time              `json:"updated_at"`
	}

	ProductVariationInfo struct {
		ID              string  `json:"id"`
		ColorType       string  `json:"color_type"`
		SizeType        string  `json:"size_type"`
		VariationPrice  float64 `json:"variation_price"`
		Quantity        int     `json:"quantity"`
		ImageVariations string  `json:"image_variations"`
		// Images          []ImageInfo `json:"images"`
	}

	ImageInfo struct {
		URL       string `json:"image_url"`
		IsPrimary bool   `json:"is_primary"`
	}

	ColorInfo struct {
		ID        string `json:"id"`
		ColorType string `json:"color_type"`
	}

	SizeInfo struct {
		ID       string `json:"id"`
		SizeType string `json:"size_type"`
	}

	Paginate struct {
		Page      int64  `query:"page" validate:"required,min=1"`
        Size       int64  `query:"size" validate:"required,min=1,max=20"`
	}	
    ItemResult struct {
		ProductDetail []*ProductDetail `json:"-"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page       int64 `json:"page"`
        TotalPage int64 `json:"totalpage"`
    }

)

type FilterRequest struct {
    Gender string `query:"gender" validate:"omitempty,oneof=male female"`
	Paginate
}

type FilterRequestBySeller struct {
    Gender    string `query:"gender" validate:"omitempty,oneof=male female"`
    IsArchive *bool  `query:"is_archive" validate:"omitempty"`
	Paginate
}

