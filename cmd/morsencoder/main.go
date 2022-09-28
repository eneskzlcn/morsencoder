package main

import (
	"github.com/eneskzlcn/morsencoder/internal/config"
	"github.com/eneskzlcn/morsencoder/internal/morsencoder"
	"github.com/eneskzlcn/morsencoder/logger"
	"github.com/eneskzlcn/morsencoder/server"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	workingEnv := getEnv("DEPLOYMENT_ENVIRONMENT", "local")

	configs, err := config.LoadConfig[config.Config](".dev/", workingEnv, "yaml")

	if err != nil {
		return err
	}

	logger, err := logger.NewZapLoggerForEnv(workingEnv, 0)
	if err != nil {
		return err
	}

	service := morsencoder.NewService(logger)
	handler := morsencoder.NewHandler(service, logger)

	server := server.New([]server.Handler{
		handler,
	}, configs.Server, logger)

	return server.Start()
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
