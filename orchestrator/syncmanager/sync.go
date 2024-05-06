package syncmanager

import (
	"errors"
	"sync"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler"
)

var errConnectorsRequired = errors.New("source and target connectors are required")

type SyncConfig struct {
	Cron   string            `yaml:"cron"`
	Source connectors.Config `yaml:"source"`
	Target connectors.Config `yaml:"target"`
}

type SyncManager struct {
	logger    logger.ILogger
	factory   factory.IFactory
	scheduler scheduler.IScheduler
}

// New return syncconfig
func New(factory factory.IFactory, scheduler scheduler.IScheduler, logger logger.ILogger) *SyncManager {
	return &SyncManager{
		factory:   factory,
		logger:    logger,
		scheduler: scheduler,
	}
}

// Manage schedules the synchronization of data between the source and target connectors
// based on the provided sync configuration.
func (s *SyncManager) Manage(config *SyncConfig) error {
	sourceConnector, err := s.factory.GetConnector(config.Source)
	if err != nil {
		s.logger.Error(err, "unable to get source connector")
		return err
	}
	targetConnector, err := s.factory.GetConnector(config.Target)
	if err != nil {
		s.logger.Error(err, "unable to get target connector")
		return err
	}
	s.scheduler.ScheduleJob(config.Cron, func() {
		s.syncData(sourceConnector, targetConnector)
	})
	return nil
}

// syncData synchronizes data between the source and target connectors.
// It retrieves the list of keys from the source connector and synchronizes
// objects corresponding to these keys between the source and target connectors.
func (s *SyncManager) syncData(source connectors.Connector, target connectors.Connector) error {
	_, err := source.ListKeys(func(keys []string) {
		s.syncObjects(source, target, keys)
	})
	if err != nil {
		s.logger.Info("unable to list keys")
		return err
	}
	return nil
}

// syncObjects asynchronously syncs objects between the source and target connectors for the provided list of keys.
// It uses goroutines to concurrently synchronize multiple objects, waits for all goroutines
// to finish before returning.
func (s *SyncManager) syncObjects(source connectors.Connector, target connectors.Connector, keys []string) {
	var wg sync.WaitGroup
	wg.Add(len(keys))
	for _, key := range keys {
		go func(k string) {
			defer wg.Done()
			s.syncObject(source, target, k)
		}(key)
	}
	wg.Wait()
}

// syncObject syncs a single object between the source and target connectors.
func (s *SyncManager) syncObject(source connectors.Connector, target connectors.Connector, objectKey string) error {
	if target.Exists(objectKey) {
		s.logger.Info("skipping syncing object as it is already exists in source", "source", source.ToString(), "target", target.ToString(), "key", objectKey)
		return nil
	}
	s.logger.Info("syncing data", "source", source.ToString(), "target", target.ToString(), "key", objectKey)
	data, err := source.Get(objectKey)
	if err != nil {
		s.logger.Error(err, "unable to get the data from source")
		return err
	}
	err = target.Put(objectKey, data)
	if err != nil {
		s.logger.Error(err, "unable to create the data in target")
		return err
	}
	s.logger.Info("successfully synced data", "source", source.ToString(), "target", target.ToString(), "key", objectKey)
	return nil
}
