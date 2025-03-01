package database

import (
	"fmt"
	"globetrotter/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgres(postgresConfig *models.Postgres) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v search_path=%v sslmode=disable TimeZone=UTC", postgresConfig.Host, postgresConfig.Username, postgresConfig.Password, postgresConfig.DatabaseName, postgresConfig.Port, postgresConfig.Schema)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
