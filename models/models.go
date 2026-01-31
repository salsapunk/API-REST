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
	CreateTask = "INSERT INTO task(title, description, created_at) VALUES($1, $2, $3) RETURNING ID"
	ShowTasks  = "SELECT id, title, description, done, created_at FROM task ORDER BY id;"
	ShowTaskBI = "SELECT id, title, description, done, created_at FROM task WHERE id = $1"
	EditTask   = "UPDATE task SET $1 = $2 WHERE id = %3;"
	DeleteTask = "DELETE FROM task WHERE id = $1;"
	// OBD == ORDER BY done
	ShowTasksOBD = "SELECT id, title, description, done, created_at FROM task ORDER BY done;"
)
