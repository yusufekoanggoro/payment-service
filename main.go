package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/payment-service/internal/factory"
	infrastructure "github.com/yusufekoanggoro/payment-service/internal/infrastructure/database"
	"github.com/yusufekoanggoro/payment-service/pkg/middleware"
)

func main() {
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create data folder: ", err)
	}

	db := infrastructure.InitDB("./data/payments.db")
	defer db.Close()

	mw := middleware.NewMiddleware(db)
	modules := factory.InitAllModule(db, mw)

	router := gin.Default()

	// add config CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apiGroup := router.Group("/api")

	for _, m := range modules {
		m.RestHandler().RegisterRoutes(apiGroup)
	}

	router.Run(":8080")
}
