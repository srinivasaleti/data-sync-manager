package syncmanager

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler"
)

type SyncManager struct {
	logger    logger.ILogger
	factory   factory.IFactory
	scheduler scheduler.IScheduler
}

func (s *SyncManager) Manage(configs []SyncConfig) {
	for _, config := range configs {
		s.scheduleSyncData(config)
	}
}

func New(factory factory.IFactory, scheduler scheduler.IScheduler, logger logger.ILogger) *SyncManager {
	return &SyncManager{
		factory:   factory,
		logger:    logger,
		scheduler: scheduler,
	}
}
