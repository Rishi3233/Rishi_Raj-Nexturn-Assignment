package main

import (
	"assignment_blog_system/config"
	"assignment_blog_system/handlers"
	"assignment_blog_system/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.InitDB()

	// Initialize Gin router
	r := gin.Default()

	// Apply middleware
	r.Use(middlewares.LogRequest)

	// Define routes
	r.POST("/blog", handlers.CreateBlog)
	r.GET("/blog/:id", handlers.GetBlog)
	r.GET("/blogs", handlers.GetBlogs)
	r.PUT("/blog/:id", handlers.UpdateBlog)
	r.DELETE("/blog/:id", handlers.DeleteBlog)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
