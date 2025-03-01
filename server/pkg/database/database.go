package database

import (
	"encoding/json"
	"fmt"
	"globetrotter/pkg/config"
	"globetrotter/pkg/logging"
	"globetrotter/pkg/models"
	"os"

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
	err = db.AutoMigrate(&models.User{}, &models.City{}, &models.Clue{}, &models.Fact{}, &models.Game{})
	if err != nil {
		logging.Error("failed in migration", "error", err)
		return nil, err
	}

	return db, nil
}

func Get() *gorm.DB {
	return db
}

func ReloadDataset(filePath string) error {

	db.Exec("truncate cities;")
	db.Exec("truncate clues;")
	db.Exec("truncate facts;")

	type data struct {
		City    string   `json:"city"`
		Country string   `json:"country"`
		Clues   []string `json:"clues"`
		Facts   []string `json:"facts"`
	}

	var dataset []data

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read dataset.json: %w", err)
	}

	err = json.Unmarshal(fileData, &dataset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal dataset.json: %w", err)
	}

	for _, record := range dataset {
		city := &models.City{
			City:    record.City,
			Country: record.Country,
		}

		db.Create(city)

		for _, clue := range record.Clues {
			db.Create(&models.Clue{
				CitiesID: city.ID,
				Clue:     clue,
			})
		}

		for _, fact := range record.Facts {
			db.Create(&models.Fact{
				CitiesID: city.ID,
				Fact:     fact,
			})
		}
	}

	return nil
}
