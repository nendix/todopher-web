package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nendix/toweb/db"
	"github.com/nendix/toweb/models"
)

// GetIndex renders the main index page with the form and todo list
func GetIndex(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos ORDER BY due_date ASC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos}) // Serve the main page
}

// GetTodos returns the list of todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos ORDER BY due_date ASC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos}) // Serve the todo list only
}

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	dueDateStr := c.PostForm("due_date")

	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD.")
		return
	}

	// Insert the new todo into the database
	_, err = db.DB.Exec("INSERT INTO todos (title, due_date) VALUES (?, ?)", title, dueDate)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating todo")
		return
	}

	// Fetch the updated todo list
	var todos []models.Todo
	err = db.DB.Select(&todos, "SELECT * FROM todos ORDER BY due_date ASC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}

	// Respond with the updated todo list HTML
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos})
}

// UpdateTodoStatus toggles the completion status of a todo
func UpdateTodoStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.DB.Exec("UPDATE todos SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating todo status")
		return
	}

	// Fetch the updated todo
	var updatedTodo models.Todo
	err = db.DB.Get(&updatedTodo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching updated todo")
		return
	}

	// Respond with the updated todo item HTML
	c.HTML(http.StatusOK, "todo_item.html", updatedTodo)
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting todo")
		return
	}

	// Respond with an empty list item (or an updated list if preferred)
	c.String(http.StatusOK, "") // Respond with no content to remove the item
}
