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
	err := db.DB.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos: %v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", nil) // Serve the main page
}

// GetTodos returns the list of todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todos")
		return
	}
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos}) // Serve the todo list only
}

func GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	err := db.DB.Get(&todo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching todo")
		return
	}
	c.HTML(http.StatusOK, "todo_item.html", todo)
}

func GetCompletedTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos WHERE completed = true")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching completed todos")
		return
	}
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos})
}

func GetPendingTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos WHERE completed = false")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching pending todos")
		return
	}
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos})
}

func GetSearchForm(c *gin.Context) {
	c.HTML(http.StatusOK, "search_form.html", nil)
}

func SearchTodos(c *gin.Context) {
	query := c.Query("query")
	var todos []models.Todo

	// Modify the SQL query to filter based on the search term
	searchQuery := "%" + query + "%"
	err := db.DB.Select(&todos, "SELECT * FROM todos WHERE title LIKE ? ORDER BY due_date ASC", searchQuery)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error searching todos")
		return
	}

	// Render only the todo list with matched results
	c.HTML(http.StatusOK, "todo_list.html", gin.H{"todos": todos})
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
	result, err := db.DB.Exec("INSERT INTO todos (title, due_date) VALUES (?, ?)", title, dueDate)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating todo")
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving new todo ID")
		return
	}
	c.Redirect(http.StatusFound, "/todos/"+strconv.Itoa(int(id)))
}

func GetEditTodo(c *gin.Context) {
	id := c.Param("id")

	// Fetch the todo from the database based on the ID
	var todo models.Todo
	err := db.DB.Get(&todo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}

	// Render the form with the todo details
	c.HTML(http.StatusOK, "edit_form.html", gin.H{"todo": todo})
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	dueDateStr := c.PostForm("due_date")

	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD.")
		return
	}

	_, err = db.DB.Exec("UPDATE todos SET title = ?, due_date = ? WHERE id = ?", title, dueDate, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating todo")
		return
	}

	// Fetch the updated todo from the database based on the ID
	var updatedTodo models.Todo
	err = db.DB.Get(&updatedTodo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}

	c.HTML(http.StatusOK, "todo_item.html", updatedTodo)
}

// UpdateTodoStatus toggles the completion status of a todo
func UpdateTodoStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.DB.Exec("UPDATE todos SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating todo status")
		return
	}

	// Fetch the todo from the database based on the ID
	var todo models.Todo
	err = db.DB.Get(&todo, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}
	// Respond with the updated todo item HTML
	c.HTML(http.StatusOK, "todo_item.html", todo)
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
