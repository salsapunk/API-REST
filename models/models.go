package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

const (
	// task with description
	CreateTaskDSQL = "INSERT INTO task(title, description, created_at) VALUES($1, $2, $3) RETURNING ID"
	ShowTasksSQL   = "SELECT id, title, description, done, created_at FROM task;"
)
