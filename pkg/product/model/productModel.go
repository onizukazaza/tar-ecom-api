package model

import "time"

type ImageInfo struct {
	URL       string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}


type ColorInfo struct {
	ID        string `json:"id"`
	ColorType string `json:"color_type"`
}

type SizeInfo struct {
	ID       string `json:"id"`
	SizeType string `json:"size_type"`
}

type ProductCreatingReq struct {
	ProductName      string                        `json:"product_name" validate:"required,max=64"`
	Description      string                        `json:"description" validate:"required,max=128"`
	SellerID         string                        `json:"seller_id" validate:"required,uuid4"`
	Gender           string                        `json:"gender" validate:"required,oneof=male female na"`
	PrimaryImage     ImageInfo                     `json:"primary_image" validate:"required"`
	AdditionalImages []ImageInfo                   `json:"additional_images,omitempty"`
	Variations       []ProductVariationCreatingReq `json:"variations,omitempty"`
	CreatedAt        time.Time                     `json:"created_at"`
	UpdatedAt        time.Time                     `json:"updated_at"`
}

type ProductVariationCreatingReq struct {
	ColorID         string  `json:"color_id" validate:"required,uuid4"`
	SizeID          string  `json:"size_id" validate:"required,uuid4"`
	VariationPrice  float64 `json:"variation_price" validate:"required,gt=0"`
	Quantity        int     `json:"quantity" validate:"required,gt=0"`
	ImageVariations string  `json:"image_variations,omitempty"`
}

type ProductEditingReq struct {
	ID               string                        `json:"id" validate:"required,uuid4"`
	ProductName      string                        `json:"product_name" validate:"omitempty,max=64"`
	Description      string                        `json:"description" validate:"omitempty,max=128"`
	Gender           string                        `json:"gender" validate:"omitempty,oneof=male female na"`
	AdditionalImages []ProductImageUpdatingReq     `json:"additional_images,omitempty"`
	Variations       []ProductVariationUpdatingReq `json:"variations,omitempty"`
}

type ProductImageUpdatingReq struct {
	ID        string `json:"id" validate:"required,uuid4"`
	ImageURL  string `json:"image_url,omitempty"`
	IsPrimary bool   `json:"is_primary"`
}


type ProductVariationUpdatingReq struct {
	ID              string  `json:"id" validate:"required,uuid4"`
	ColorID         string  `json:"color_id,omitempty"`
	SizeID          string  `json:"size_id,omitempty"`
	VariationPrice  float64 `json:"variation_price,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	ImageVariations string  `json:"image_variations,omitempty"`
}
