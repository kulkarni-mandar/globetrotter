package controller

import (
	"encoding/json"
	"globetrotter/pkg/config"
	"globetrotter/pkg/database"
	"globetrotter/pkg/logging"
	"globetrotter/pkg/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func healthController(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"status": "running"},
	)
}

func configController(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		config.Get(),
	)
}

func refreshDataset(c *gin.Context) {
	db := database.Get()

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

	fileData, err := os.ReadFile("dataset.json")
	if err != nil {
		logging.Error("failed to read dataset.json", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to find dataset",
		})

		return
	}

	err = json.Unmarshal(fileData, &dataset)
	if err != nil {
		logging.Error("failed to unmarshal dataset.json", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load dataset",
		})

		return
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
