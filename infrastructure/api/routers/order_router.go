package routers

import (
	"hub-poc-api-v2/infrastructure/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type OrderRouter struct {
	group        *gin.RouterGroup
	databasePool *sqlx.DB
}

func (r *OrderRouter) Route() {
	handler := &handlers.OrderHandler{}

	r.group.GET("orders", handler.GetOrders)
	r.group.GET("orders/:id", handler.GetOrder)
}
