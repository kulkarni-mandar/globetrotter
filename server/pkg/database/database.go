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
	default:
		db = nil
		err = fmt.Errorf("invalid database type: %v", dbConfig.Type)
	}

	if err != nil {
		return nil, err
	}

	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %v", config.Get().Server.Name))

	return db, nil
}

func Get() *gorm.DB {
	return db
}
