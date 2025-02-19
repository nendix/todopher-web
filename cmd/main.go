package main

import (
	"log"
	"net/http"

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
	r.Static("/todopher/static", "web/static")

	// Gruppo di rotte con prefisso /todopher
	todo := r.Group("/todopher")
	{

		// Route for the main index page
		todo.GET("/", handlers.GetIndex)

		// API routes for todos
		todo.GET("/todos", handlers.GetTodos)
		todo.GET("/todos/pending", handlers.GetPendingTodos)
		todo.GET("/todos/completed", handlers.GetCompletedTodos)
		todo.GET("/todos/:id", handlers.GetTodoByID)
		todo.GET("/todos/search", handlers.SearchTodos)
		todo.GET("/search-form", handlers.GetSearchForm)
		todo.GET("/todos/edit/:id", handlers.GetEditTodo)
		todo.POST("/todos", handlers.CreateTodo)
		todo.PATCH("/todos/:id", handlers.UpdateTodo)
		todo.PATCH("/todos/:id/toggle", handlers.UpdateTodoStatus)
		todo.DELETE("/todos/:id", handlers.DeleteTodo)
	}

	// Start server
	r.SetTrustedProxies(nil)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
