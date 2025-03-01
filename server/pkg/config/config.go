package config

import (
	"globetrotter/pkg/logging"
	"globetrotter/pkg/models"

	"github.com/spf13/viper"
)

var config *models.Config

func init() {
	config = &models.Config{}
}

func New(configPath *string) *models.Config {
	viper.SetConfigFile(*configPath)

	if err := viper.ReadInConfig(); err != nil {
		logging.Fatal(
			"failed to read config",
			"configPath", *configPath,
		)
	}

	config = &models.Config{
		Name: viper.GetString("server.name"),
		Port: viper.GetString("server.port"),
	}

	return config
}

func Get() *models.Config {
	return config
}
