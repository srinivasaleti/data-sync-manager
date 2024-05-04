package main

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("starting server")
	s := syncmanager.New(factory.New(), logger)
	s.Manage([]syncmanager.SyncConfig{
		{
			Source: "s3",
			Target: "local",
		},
	})
	logger.Info("shutting down the server")
	logger.Info("successfully shutdown the server")
}
