package db

import (
	"fmt"
	// "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() error {
	dbUser := "user"      // MySQL username
	dbPassword := "pass"  // MySQL password
	dbHost := "localhost" // MySQL container name or host IP
	dbPort := "3306"      // Default MySQL port
	dbName := "tododb"    // Database name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	DB = db

	// Create the todos table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		due_date DATE NOT NULL
	);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	return nil
}
