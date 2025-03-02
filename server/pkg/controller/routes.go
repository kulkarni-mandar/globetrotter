package controller

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.RouterGroup) {

	internal := router.Group("/internal")
	{
		internal.GET("/health", healthController)
		internal.GET("/config", configController)
		internal.POST("/dataset/refresh", refreshDataset)
	}

	users := router.Group("/users")
	{
		users.POST("", addNewUser)
	}

	game := router.Group("/game")
	{
		game.POST("/play/:userName", newOrJoinGame)
		game.POST("/play/:userName/:sessionId/invite", inviteToGame)
		game.POST("/play/:userName/:sessionId/next", nextQuestion)
		game.POST("/play/:userName/:sessionId/:clueId/validate", validateAnswer)
		game.POST("/play/:userName/:sessionId/end", endGame)
	}

}
