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

	// Routes
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route is working")
	})

	// Route for the main index page
	r.GET("/todopher", handlers.GetIndex)

	// API routes for todos
	r.GET("/todopher/todos", handlers.GetTodos)                      // Fetch only the todo list
	r.GET("/todopher/todos/pending", handlers.GetPendingTodos)       // Fetch only the pending todo list
	r.GET("/todopher/todos/completed", handlers.GetCompletedTodos)   // Fetch only the completed todo list
	r.GET("/todopher/todos/:id", handlers.GetTodoByID)               // Fetch a single todo
	r.GET("/todopher/todos/search", handlers.SearchTodos)            // Search for todos
	r.GET("/todopher/search-form", handlers.GetSearchForm)           // Search for todos
	r.GET("/todopher/todos/edit/:id", handlers.GetEditTodo)          // Load todo data in form
	r.POST("/todopher/todos", handlers.CreateTodo)                   // Create a new todo
	r.PATCH("/todopher/todos/:id", handlers.UpdateTodo)              // Edit a todo
	r.PATCH("/todopher/todos/:id/toggle", handlers.UpdateTodoStatus) // Toggle todo status
	r.DELETE("/todopher/todos/:id", handlers.DeleteTodo)             // Delete a todo

	// Start server
	r.SetTrustedProxies(nil)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
