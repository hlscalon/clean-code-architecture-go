package main

import (
	"hub-poc-api-v2/infrastructure/api/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	pool, err := sqlx.Connect("postgres", "user=foo dbname=bar")

	if err != nil {
		log.Fatal(err)
		return
	}

	gin := gin.Default()
	routers.Setup(gin, pool)

	gin.Run("localhost:8080")
}
