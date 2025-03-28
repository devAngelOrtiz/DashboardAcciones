package main

import (
	"context"
	"log"

	"api/common"
	"api/scrapers"
	"api/stock"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error al cargar .env: %v", err)
	}

	conn, err := common.GetDBConnection()
	if err != nil {
		log.Fatalf("❌ Error al conectar a la base de datos: %v", err)
	}
	defer conn.Close(context.Background())

	err = scrapers.ScrapeAndStoreData(conn)
	if err != nil {
		log.Fatalf("❌ Error al hacer scraping y almacenar los datos: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/api")

	stock.StockRegister(v1.Group("/stock"), conn)

	r.Run(":3000")
}
