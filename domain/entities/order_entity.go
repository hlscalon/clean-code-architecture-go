package entities

import "time"

type OrderEntity struct {
	Id          int
	CustomerId  int
	CreatedDate time.Time
	TotalPrice  float64
}
