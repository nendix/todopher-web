package main

import (
	"log"
	"net/http"

	"github.com/nendix/toweb/db"
	"github.com/nendix/toweb/handlers"

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

	// Routes
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route is working")
	})

	// Route for the main index page
	r.GET("/", handlers.GetIndex)

	// API routes for todos
	r.GET("/todos", handlers.GetTodos)             // Fetch only the todo list
	r.GET("/todos/:id", handlers.GetTodoByID)      // Fetch a single todo
	r.GET("/todos/edit/:id", handlers.GetEditTodo) // Load todo data in form
	r.POST("/todos", handlers.CreateTodo)          // Create a new todo
	r.PATCH("/todos/:id", handlers.UpdateTodo)     // Edit a todo
	// r.PATCH("/todos/:id/toggle", handlers.UpdateTodoStatus) // Toggle todo status
	r.DELETE("/todos/:id", handlers.DeleteTodo) // Delete a todo

	// Start server
	r.SetTrustedProxies(nil)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
