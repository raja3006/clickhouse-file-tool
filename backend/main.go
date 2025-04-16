package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rajaverma/clickhouse-file-tool/pkg/api"
	"github.com/rajaverma/clickhouse-file-tool/pkg/clickhouse"
	"github.com/rajaverma/clickhouse-file-tool/pkg/file"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Initialize handlers
	handlers := api.NewHandlers(nil, nil)

	// API routes
	api := r.Group("/api")
	{
		// ClickHouse routes
		api.POST("/clickhouse/connect", handlers.ConnectClickHouse)
		api.GET("/clickhouse/tables", handlers.GetTables)
		api.GET("/clickhouse/columns/:table", handlers.GetColumns)

		// File routes
		api.POST("/file/columns", handlers.GetFileColumns)

		// Ingestion routes
		api.POST("/ingest/:source/:target", handlers.IngestData)
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 