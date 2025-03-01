package controller

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

	internal := router.Group("/internal")
	{
		internal.GET("/health", healthController)
	}

}
