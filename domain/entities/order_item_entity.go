package entities

type OrderItemEntity struct {
	Id        int
	OrderId   int
	ProductId int
	Price     float64
	Quantity  int
}
