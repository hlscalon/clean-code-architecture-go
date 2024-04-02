package main

import (
	"fmt"
	"hub-poc-api-v2/infrastructure/api/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connectDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	return sqlx.Connect("postgres", dsn)
}

func main() {
	pool, err := connectDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	gin := gin.Default()
	routers.Setup(gin, pool)

	gin.Run("localhost:8080")
}
