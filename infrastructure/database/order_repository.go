package database

import (
	"hub-poc-api-v2/application/usecases/orders"
	"hub-poc-api-v2/domain/entities"
	"log"

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
			"FROM \"order\" "+
			"WHERE id = $1",
		id,
	)

	if err != nil {
		log.Printf("Erro: %v\n", err)
		return entities.OrderEntity{}, err
	}

	return orderEntity, nil
}

func (r *OrderRepository) GetOrdersWithFilter(filter *orders.GetOrdersFilter) ([]entities.OrderEntity, error) {
	orders := []entities.OrderEntity{}

	query :=
		"SELECT * " +
			"FROM \"order\" "

	var err error

	if filter.FromDate != "" && filter.ToDate != "" {
		query += "WHERE created_date >= $1 " +
			"AND created_date <= $2 "

		err = r.DatabasePool.Select(
			&orders,
			query,
			filter.FromDate, filter.ToDate,
		)
	} else {
		err = r.DatabasePool.Select(&orders, query)
	}

	if err != nil {
		log.Printf("Erro: %v\n", err)
		return nil, err
	}

	return orders, nil
}
