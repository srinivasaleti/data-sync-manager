package syncmanager

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
)

type SyncManager struct {
	logger  logger.ILogger
	factory factory.IFactory
}

func (s *SyncManager) Manage(configs []SyncConfig) {
	for _, config := range configs {
		s.scheduleSyncData(config)
	}
}

func New(factory factory.IFactory, logger logger.ILogger) *SyncManager {
	return &SyncManager{
		factory: factory,
		logger:  logger,
	}
}
