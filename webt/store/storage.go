package store

import "database/sql"

type Store interface {
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}
