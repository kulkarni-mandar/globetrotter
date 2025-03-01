package main

import (
	"flag"
	"globetrotter/pkg/config"
	"globetrotter/pkg/controller"
	"globetrotter/pkg/logging"
	"globetrotter/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	// read config path
	flag.StringVar(&configPath, "config", "./app.yaml", "config for server")
	flag.Parse()

	logging.Info("config path parsed", "config", configPath)

	// init config
	config.New(&configPath)

	logging.Debug("loaded config", "config", config.Get())
}

func main() {

	router := gin.New()

	router.Use(
		middlewares.Logger(),
		middlewares.PanicRecovery(),
	)

	controller.SetupRoutes(router.Group("/api"))

	logging.Info(
		"starting server",
		"port", config.Get().Server.Port,
	)

	router.Run(":" + config.Get().Server.Port)
}
