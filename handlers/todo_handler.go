// handlers/todo_handler.go
package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/nendix/toweb/db"
	"github.com/nendix/toweb/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos ORDER BY due_date ASC")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}

func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	dueDateStr := c.PostForm("due_date")

	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD.")
		return
	}

	_, err = db.DB.Exec("INSERT INTO todos (title, due_date) VALUES (?, ?)", title, dueDate)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating todo")
		return
	}
	GetTodos(c) // Respond with the updated list
}

func UpdateTodoStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.DB.Exec("UPDATE todos SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating todo status")
		return
	}
	GetTodos(c) // Respond with the updated list
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting todo")
		return
	}
	GetTodos(c) // Respond with the updated list
}
