package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(gin *gin.Engine, databasePool *sqlx.DB) {
	router := gin.Group("")

	// Inclui todos os middlewares necessários
	// router.Use(middlewares...)

	orderRouter := &OrderRouter{
		group:        router,
		databasePool: databasePool,
	}
	orderRouter.Route()
}
