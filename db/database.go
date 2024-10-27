package db

import (
	"fmt"
	// "os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() error {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	dsn := "todo_user:pass@tcp(db:3306)/tododb?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	DB = db

	// Create the todos table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		due_date DATE
	);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	// Check if the table is empty, and add a test to-do if it is
	var count int
	err = DB.Get(&count, "SELECT COUNT(*) FROM todos")
	if err != nil {
		return fmt.Errorf("error counting todos: %w", err)
	}

	if count == 0 {
		// Insert a test to-do item
		_, err = DB.Exec("INSERT INTO todos (title, completed, due_date) VALUES (?, ?, ?)",
			"Test To-Do Item", false, time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))
		if err != nil {
			return fmt.Errorf("error inserting test todo: %w", err)
		}
	}

	return nil
}
