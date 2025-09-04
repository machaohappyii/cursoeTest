package routes

import (
	"gin-user-api/controllers"
	"gin-user-api/middleware"
	"gin-user-api/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create Gin router
	router := gin.New()
	
	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.AccessTimer())
	
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
			"message": "Server is running",
		})
	})
	
	// Test endpoint for AccessTimer middleware
	router.GET("/test", func(c *gin.Context) {
		// Simulate some processing time
		time.Sleep(100 * time.Millisecond)
		c.JSON(200, gin.H{
			"message": "Test endpoint - processing complete",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})
	
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		setupAuthRoutes(v1, db)
		setupUserRoutes(v1, db)
	}
	
	return router
}

func setupAuthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)
	
	auth := rg.Group("/auth")
	{
		auth.POST("/register", userController.CreateUser)
		auth.POST("/login", userController.Login)
	}
}

func setupUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)
	
	users := rg.Group("/users")
	users.Use(middleware.JWTAuth())
	{
		users.GET("", userController.GetUsers)
		users.GET("/:id", userController.GetUserByID)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
}