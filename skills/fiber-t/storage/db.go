package storage

import "gorm.io/gorm"

type DBs struct {
	db *gorm.DB
}

func NewDBs() *DBs {
	return &DBs{}
}

func InitStorage() error {
	return nil
}
