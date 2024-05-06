package main

import (
	"flag"
	"fmt"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/config"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"

	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	logger := logger.NewLogger()
	configFile := flag.String("config", "", "YAML file path")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("Please provide config file, Usage: --config <yaml_file_path>")
		return
	}
	config, err := config.GetConfig(*configFile)
	if err != nil {
		logger.Error(err, "unable read config")
		return
	}
	logger.Info("starting scheduler")
	ctx := ctrl.SetupSignalHandler()
	jobScheduler, err := scheduler.New()
	if err != nil {
		logger.Error(err, "unable to create scheduler")
		return
	}
	connectorsFactory := factory.New(logger)
	s := syncmanager.New(connectorsFactory, jobScheduler, logger)
	s.Manage(config)
	jobScheduler.Start()
	<-ctx.Done()
	logger.Info("shutting down scheduler")
	jobScheduler.Shutdown()
	logger.Info("successfully shutdown the server")
}
