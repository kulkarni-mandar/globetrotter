package main

import (
	"flag"
	"globetrotter/pkg/config"
	"globetrotter/pkg/controller"
	"globetrotter/pkg/database"
	"globetrotter/pkg/logging"
	"globetrotter/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var configPath string
var initDatasetPath string

func init() {
	// read config path
	flag.StringVar(&configPath, "config", "./app.yaml", "config for server")
	flag.StringVar(&initDatasetPath, "dataset", "./dataset.json", "dataset for initializing the quiz")

	flag.Parse()

	logging.Info("config path parsed", "config", configPath)
	logging.Info("dataset path parsed", "dataset", initDatasetPath)

	// init config
	config.New(&configPath)

	database.New()

	logging.Debug("connected to database")

	err := database.ReloadDataset(initDatasetPath)
	if err != nil {
		logging.Error("error in loading dataset", "error", err)
	}
}

func main() {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(
		middlewares.Cors(),
		middlewares.Logger(),
		middlewares.PanicRecovery(),
	)

	controller.SetupRoutes(router.Group("/api"))

	logging.Info(
		"starting server",
		"port", config.Get().Server.Port,
	)

	router.Run("0.0.0.0:" + config.Get().Server.Port)
}
