package orders

import "hub-poc-api-v2/domain/entities"

type GetOrderUseCase struct {
	OrderRepository OrderRepository
}

func (uc *GetOrderUseCase) Execute(id int) (entities.OrderEntity, error) {
	return uc.OrderRepository.GetOrder(id)
}
