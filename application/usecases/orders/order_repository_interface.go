package orders

import "hub-poc-api-v2/domain/entities"

type OrderRepository interface {
	GetOrder(id int) (entities.OrderEntity, error)
	GetOrdersWithFilter(fromDate, toDate string) ([]entities.OrderEntity, error)
}
