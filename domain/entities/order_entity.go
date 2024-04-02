package entities

import "time"

type OrderEntity struct {
	Id          int       `json:"id"`
	CustomerId  int       `json:"customer_id" db:"customer_id"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	TotalPrice  float64   `json:"total_price" db:"total_price"`
}
