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
		ID              string      `json:"id"`
		ColorType       string      `json:"color_type"`
		SizeType        string      `json:"size_type"`
		VariationPrice  float64     `json:"variation_price"`
		Quantity        int         `json:"quantity"`
		ImageVariations string      `json:"image_variations"`
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

	
//  not use !!!!
	ProductCreatingReq struct {
		ProductName      string                        `json:"product_name" validate:"required,max=64"`
		Description      string                        `json:"description" validate:"required,max=128"`
		SellerID         string                        `json:"seller_id" validate:"required,uuid4"`
		Gender           string                        `json:"gender" validate:"required,oneof=male female na"`
		PrimaryImage     ImageInfo                     `json:"primary_image" validate:"required"`
		AdditionalImages []ImageInfo                   `json:"additional_images,omitempty"`
		Variations       []ProductVariationCreatingReq `json:"variations,omitempty"`
	}


	ProductVariationCreatingReq struct {
		ColorID         string  `json:"color_id" validate:"required,uuid4"`
		SizeID          string  `json:"size_id" validate:"required,uuid4"`
		VariationPrice  float64 `json:"variation_price" validate:"required,min=0"`
		Quantity        int     `json:"quantity" validate:"required,min=0"`
		ImageVariations string  `json:"image_variations,omitempty"`
	}

	ProductEditingReq struct {
		ID               string                        `json:"id" validate:"required,uuid4"`
		ProductName      string                        `json:"product_name" validate:"omitempty,max=64"`
		Description      string                        `json:"description" validate:"omitempty,max=128"`
		SellerID         string                        `json:"seller_id" validate:"omitempty,uuid4"`
		Gender           string                        `json:"gender" validate:"omitempty,oneof=male female na"`
		PrimaryImage     ImageInfo                     `json:"primary_image" validate:"omitempty"`
		AdditionalImages []ImageInfo                   `json:"additional_images,omitempty"`
		Variations       []ProductVariationCreatingReq `json:"variations,omitempty"`
	}
	
	Paginate struct {
        Page      int64 `query:"page" validate:"omitempty,min=1"`
        Size      int64 `query:"size" validate:"omitempty,min=1"`
    }

    PaginateResult struct {
        Page      int64 `json:"page"`
        TotalPage int64 `json:"total_page"`
    }


    ProductResult struct {
        Items    []*ProductDetail `json:"items"`
        Paginate PaginateResult  `json:"paginate"`
    }

	FilterRequest struct {
			Gender string `query:"gender" validate:"omitempty,oneof=male female other"`
			IsArchive bool   `query:"is_archive" validate:"omitempty"`
			Paginate   
	}
)
