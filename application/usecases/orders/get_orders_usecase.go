package orders

import "hub-poc-api-v2/domain/entities"

type GetOrdersUseCase struct {
	OrderRepository OrderRepository
}

func (uc *GetOrdersUseCase) Execute(filter *GetOrdersFilter) ([]entities.OrderEntity, error) {
	return uc.OrderRepository.GetOrdersWithFilter(filter)
}
