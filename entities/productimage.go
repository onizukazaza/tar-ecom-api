package entities

import "time"

type ProductImage struct {
	ID        int       `db:"id"`         
	ImageUrl  string    `db:"image_url"`  
	ProductID int       `db:"product_id"` 
	IsPrimary bool      `db:"is_primary"` 
	CreatedAt time.Time `db:"created_at"` 
	UpdatedAt time.Time `db:"updated_at"` 
}
