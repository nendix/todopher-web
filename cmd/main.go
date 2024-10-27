package main

import (
	"log"

	"github.com/nendix/toweb/db"
	"github.com/nendix/toweb/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	if err := db.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*.html")

	// Routes
	r.GET("/", handlers.GetTodos)
	r.POST("/todos", handlers.CreateTodo)
	r.POST("/todos/:id/toggle", handlers.UpdateTodoStatus)
	r.DELETE("/todos/:id", handlers.DeleteTodo)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
