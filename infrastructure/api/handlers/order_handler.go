package handlers

import (
	"hub-poc-api-v2/adapters/controllers/orders"
	uc_orders "hub-poc-api-v2/application/usecases/orders"
	"hub-poc-api-v2/infrastructure/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type OrderHandler struct {
	DatabasePool *sqlx.DB
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	filter := &uc_orders.GetOrdersFilter{
		FromDate: c.Query("from_date"),
		ToDate:   c.Query("to_date"),
	}

	controller := &orders.GetOrdersController{
		GetOrdersUseCase: &uc_orders.GetOrdersUseCase{
			OrderRepository: &database.OrderRepository{
				DatabasePool: h.DatabasePool,
			},
		},
	}
	orders, err := controller.Execute(filter)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	controller := &orders.GetOrderController{
		GetOrderUseCase: &uc_orders.GetOrderUseCase{
			OrderRepository: &database.OrderRepository{
				DatabasePool: h.DatabasePool,
			},
		},
	}
	order, err := controller.Execute(
		id,
	)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, order)
}
