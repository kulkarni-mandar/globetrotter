package controller

import (
	"globetrotter/pkg/logging"
	"globetrotter/pkg/models"
	"globetrotter/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addNewUser(c *gin.Context) {
	user := models.User{}

	err := service.AddNewUser(&user)
	if err != nil {
		logging.Error("failed to create new user", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to crete new user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
