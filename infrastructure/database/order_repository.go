package database

import (
	"hub-poc-api-v2/application/usecases/orders"
	"hub-poc-api-v2/domain/entities"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	DatabasePool *sqlx.DB
}

func (r *OrderRepository) GetOrder(id int) (entities.OrderEntity, error) {
	orderEntity := entities.OrderEntity{}
	err := r.DatabasePool.Get(
		&orderEntity,
		"SELECT * "+
			"FROM order "+
			"WHERE id = $1",
		id,
	)

	if err != nil {
		return entities.OrderEntity{}, err
	}

	return orderEntity, nil
}

func (r *OrderRepository) GetOrdersWithFilter(filter *orders.GetOrdersFilter) ([]entities.OrderEntity, error) {
	orders := []entities.OrderEntity{}

	err := r.DatabasePool.Select(
		&orders,
		"SELECT * "+
			"FROM order "+
			"WHERE created_date >= $1 "+
			"AND created_date <= $2 ",
		filter.FromDate, filter.ToDate,
	)

	if err != nil {
		return nil, err
	}

	return orders, nil
}
