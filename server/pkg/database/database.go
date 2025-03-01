package database

import (
	"fmt"
	"globetrotter/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

func New() (*gorm.DB, error) {
	dbConfig := config.Get().Database

	var err error

	switch dbConfig.Type {
	case "postgres":
		db, err = newPostgres(dbConfig.Postgres)
		return db, err
	default:
		return nil, fmt.Errorf("invalid database type: %v", dbConfig.Type)
	}
}

func Get() *gorm.DB {
	return db
}
