package handlers

import (
	"hub-poc-api-v2/adapters/controllers/orders"
	uc_orders "hub-poc-api-v2/application/usecases/orders"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type OrderHandler struct {
	DatabasePool *sqlx.DB
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	controller := &orders.GetOrdersController{}

	filter := &uc_orders.GetOrdersFilter{
		FromDate: c.Query("from_date"),
		ToDate:   c.Query("to_date"),
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

	controller := &orders.GetOrderController{}
	order, err := controller.Execute(
		id,
	)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, order)
}
