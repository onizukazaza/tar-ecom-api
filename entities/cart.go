package entities

import "time"

type Cart struct {
	ID         int       `db:"id"`          
	UserID     int       `db:"user_id"`     
	TotalPrice uint      `db:"total_price"`
	CreatedAt  time.Time `db:"created_at"`  
	UpdatedAt  time.Time `db:"updated_at"`  
}
