package main

import (
	"log"

	"github.com/nendix/todopher-web/db"
	"github.com/nendix/todopher-web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	if err := db.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.DB.Close()

	// Initialize Gin router
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "web/static")

	// Gruppo di rotte con prefisso /todopher

	// Route for the main index page
	r.GET("/", handlers.GetIndex)

	// API routes for todos
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/pending", handlers.GetPendingTodos)
	r.GET("/todos/completed", handlers.GetCompletedTodos)
	r.GET("/todos/:id", handlers.GetTodoByID)
	r.GET("/todos/search", handlers.SearchTodos)
	r.GET("/search-form", handlers.GetSearchForm)
	r.GET("/todos/edit/:id", handlers.GetEditTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PATCH("/todos/:id", handlers.UpdateTodo)
	r.PATCH("/todos/:id/toggle", handlers.UpdateTodoStatus)
	r.DELETE("/todos/:id", handlers.DeleteTodo)

	// Start server
	r.SetTrustedProxies(nil)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
