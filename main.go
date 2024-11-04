//go:build !test
// +build !test

package main

import (
	"go_api_template/api"
	"go_api_template/pkg/utils"

	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger := utils.NewLogger()
	config := utils.NewConfig(".env")
	db := utils.NewPgPool(logger, config)

	defer db.Close()

	if config.Port == "" {
		logger.Error("No port specified in the environment, defaulting to 8080")
		config.Port = "8080"
	}

	if err := api.Run(ctx, logger, config, db); err != nil {
		logger.Errorf("Server stopped with error: %v\n", err)
	}
}
