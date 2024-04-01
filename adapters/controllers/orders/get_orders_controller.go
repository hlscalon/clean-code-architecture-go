package orders

import (
	"hub-poc-api-v2/application/usecases/orders"
	"hub-poc-api-v2/domain/entities"
)

type GetOrdersController struct {
	GetOrdersUseCase orders.GetOrdersUseCase
}

func (c *GetOrdersController) Execute(fromDate, toDate string) ([]entities.OrderEntity, error) {
	return c.GetOrdersUseCase.Execute(fromDate, toDate)
}
