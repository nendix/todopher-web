package models

import "time"

type Todo struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Completed bool      `db:"completed" json:"completed"`
	DueDate   time.Time `db:"due_date" json:"due_date"`
}
