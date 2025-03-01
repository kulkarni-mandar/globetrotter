package controller

import (
	"globetrotter/pkg/config"
	"globetrotter/pkg/database"
	"net/http"

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
	err := database.ReloadDataset("dataset.json")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to reload dataset",
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
