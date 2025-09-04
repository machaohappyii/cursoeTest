package main

import (
	"gin-user-api/config"
	"gin-user-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	
	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)
	
	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	
	// Setup router
	router := routes.SetupRouter(db)
	
	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(cfg.GetServerAddr()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}