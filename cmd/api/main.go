package main

import (
	"go-clean-architecture/config"
	"go-clean-architecture/pkg/utils/logger"
	"go-clean-architecture/server"
)

func main() {
	// Init utils and configs
	logger.InitLogger()
	appConfig, err := config.Init()

	if err != nil {
		logger.GetLogger().Fatalf("%s", err.Error())
	}

	app := server.NewApp(appConfig)
	if err := app.Run(appConfig.Server.Port); err != nil {
		logger.GetLogger().Fatalf("%s", err.Error())
	}
}
