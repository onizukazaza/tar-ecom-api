package entities

import "time"

type ProductVariation struct {
	ID             int       `db:"id"`                 
	ProductID      int       `db:"product_id"`         
	VariationValue string    `db:"variation_value"`    
	ImageVariation string    `db:"image_variation"`    
	Price          uint      `db:"price"`              
	CreatedAt      time.Time `db:"created_at"`         
	UpdatedAt      time.Time `db:"updated_at"`         
}
