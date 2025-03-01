package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthController(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"status": "running"},
	)
}
