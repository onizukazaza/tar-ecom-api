package entities

import "time"

type Product struct {
	ID               int       `db:"id"`                
	SellerID         int       `db:"seller_id"`         
	ProductName      string    `db:"product_name"`      
	Description      string    `db:"description"`       
	IsArchive        bool      `db:"is_archive"`        
	CreatedAt        time.Time `db:"created_at"`        
	UpdatedAt        time.Time `db:"updated_at"`        
}
