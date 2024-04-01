package orders

import (
	"hub-poc-api-v2/application/usecases/orders"
	"hub-poc-api-v2/domain/entities"
)

type GetOrderController struct {
	GetOrderUseCase orders.GetOrderUseCase
}

func (c *GetOrderController) Execute(id int) (entities.OrderEntity, error) {
	return c.GetOrderUseCase.Execute(id)
}
