package main

// CREATE TABLE IF NOT EXISTS tasks (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	project_id INT,
// 	name VARCHAR(255) NOT NULL,
// 	description TEXT,
// 	status VARCHAR(255) NOT NULL,
// 	FOREIGN KEY (project_id) REFERENCES projects(id)
// );
// `)

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID          int64     `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
