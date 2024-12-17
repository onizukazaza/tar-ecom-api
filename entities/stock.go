package entities

type Stock struct {
	ID                       int  `db:"id"`                         
	ProductVariationOptionID int  `db:"product_variation_option_id"` 
	Quantity                 uint `db:"quantity"`                   
}
