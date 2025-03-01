package main

import (
	"flag"
	"globetrotter/pkg/config"
	"globetrotter/pkg/logging"
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

func main() {}
