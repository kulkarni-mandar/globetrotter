package controller

import (
	"globetrotter/pkg/config"
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
