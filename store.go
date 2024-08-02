package main

import "database/sql"

type Store interface {
	// Users
	CreateUser() error
	CreateTask(t *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {

	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, description, project_id) VALUES (?, ?, ?)", t.Name, t.Description, t.ProjectID)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil

}
