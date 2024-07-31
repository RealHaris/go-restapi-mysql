package main

import "database/sql"

type Store interface {
	// Users
	CreateUser() error
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
