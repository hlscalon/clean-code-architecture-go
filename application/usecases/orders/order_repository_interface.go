package orders

import "hub-poc-api-v2/domain/entities"

type GetOrdersFilter struct {
	FromDate string
	ToDate   string
}

type OrderRepository interface {
	GetOrder(id int) (entities.OrderEntity, error)
	GetOrdersWithFilter(filter *GetOrdersFilter) ([]entities.OrderEntity, error)
}
