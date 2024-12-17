package entities

import "time"

type Address struct {
	ID            int       `db:"id"`             
	UserID        int       `db:"user_id"`         
	RecipientName string    `db:"recipient_name"`  
	PostalAddress string    `db:"postal_address"`  
	HouseNumber   string    `db:"house_number"`    
	ContactNumber string    `db:"contact_number"`  
	Favorite      bool      `db:"favorite"`        
	CreatedAt     time.Time `db:"created_at"`      
	UpdatedAt     time.Time `db:"updated_at"`      
}
