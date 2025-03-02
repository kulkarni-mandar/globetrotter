package controller

import (
	"globetrotter/pkg/logging"
	"globetrotter/pkg/models"
	"globetrotter/pkg/service"
	"globetrotter/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newOrJoinGame(c *gin.Context) {
	userName := c.Param("userName")
	sessionId := utils.ToInt(c.Query("sessionId"))

	newGame := false

	if sessionId == 0 {
		newGame = true
	}

	if newGame {
		response, err := service.NewGame(userName)
		if err != nil {
			logging.Error("failed to create new game", "error", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to create new game"})
			return
		}

		c.JSON(http.StatusOK, response)
	} else {
		response, err := service.JoinGame(userName, sessionId)
		if err != nil {
			logging.Error("failed to join game", "error", err.Error(), "sessionID", sessionId)
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to join gane"})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func inviteToGame(c *gin.Context) {
	userName := c.Param("userName")
	sessionId := utils.ToInt(c.Param("sessionId"))

	response, err := service.InviteToGame(userName, sessionId)
	if err != nil {
		logging.Error("failed to invite to game", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to invite to game"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func nextQuestion(c *gin.Context) {
	userName := c.Param("userName")
	sessionId := utils.ToInt(c.Param("sessionId"))

	response, err := service.NextQuestion(userName, sessionId)
	if err != nil {
		logging.Error("failed to move next question", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to show next question"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func validateAnswer(c *gin.Context) {
	userName := c.Param("userName")
	sessionId := utils.ToInt(c.Param("sessionId"))
	clueId := utils.ToInt(c.Param("clueId"))

	var request models.City

	err := c.BindJSON(&request)
	if err != nil {
		logging.Error("failed to read request body", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: "Invalid request body"})
		return
	}

	response, err := service.ValidateAnswer(userName, sessionId, clueId, request.ID)
	if err != nil {
		logging.Error("failed to validate answer", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to validate answer"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func endGame(c *gin.Context) {
	userName := c.Param("userName")
	sessionId := utils.ToInt(c.Param("sessionId"))

	response, err := service.EndGame(userName, sessionId)
	if err != nil {
		logging.Error("failed to finish game", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{Message: "failed to finish game"})
		return
	}

	c.JSON(http.StatusOK, response)
}
