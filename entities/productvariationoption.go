package entities

type ProductVariationOption struct {
	ID                 int    `db:"id"`                   
	ProductVariationID int    `db:"product_variation_id"` 
	OptionValue        string `db:"option_value"`        
}
