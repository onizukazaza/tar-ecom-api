package entities

type CartItem struct {
	ID                       int  `db:"id"`                          
	CartID                   int  `db:"cart_id"`                    
	ProductVariationOptionID int  `db:"product_variation_option_id"` 
	TotalPrice               uint `db:"total_price"`                 
}
