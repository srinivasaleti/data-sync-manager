package main

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"

	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("starting scheduler")
	ctx := ctrl.SetupSignalHandler()
	jobScheduler, err := scheduler.New()
	if err != nil {
		logger.Error(err, "unable to create scheduler")
		return
	}
	connectorsFactory := factory.New(logger)
	s := syncmanager.New(connectorsFactory, jobScheduler, logger)

	s.Manage([]syncmanager.SyncConfig{
		{
			Cron:      "* * * * * *",
			Source:    connectors.Config{Type: "s3"},
			ObjectKey: "20240101_062739.jpg",
			Target:    connectors.Config{Type: "local"},
		},
	})
	jobScheduler.Start()
	<-ctx.Done()
	logger.Info("shutting down scheduler")
	jobScheduler.Shutdown()
	logger.Info("successfully shutdown the server")
}
