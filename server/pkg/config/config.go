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
		Server: &models.Server{
			Name: viper.GetString("server.name"),
			Port: viper.GetString("server.port"),
		},
		Database: &models.Database{
			Type: viper.GetString("database.type"),
		},
	}

	switch config.Database.Type {

	case "postgres":
		config.Database.Postgres = &models.Postgres{
			Username:     viper.GetString("database.postgres.username"),
			Password:     viper.GetString("database.postgres.password"),
			Port:         viper.GetString("database.postgres.port"),
			DatabaseName: viper.GetString("database.postgres.dbName"),
			Host:         viper.GetString("database.postgre.host"),
		}

	}

	return config
}

func Get() *models.Config {
	return config
}
