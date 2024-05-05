package syncmanager

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

type SyncConfig struct {
	Cron      string
	Source    connectors.Config
	Target    connectors.Config
	ObjectKey string
}

func (s *SyncManager) scheduleSyncData(config SyncConfig) error {
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
		s.syncData(sourceConnector, targetConnector, config.ObjectKey)
	})
	return nil
}

func (s *SyncManager) syncData(source connectors.Connector, target connectors.Connector, objectKey string) error {
	s.logger.Info("syncing data", "source", source.ToString(), "target", target.ToString())
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
	s.logger.Info("successfully synced data", "source", source.ToString(), "target", target.ToString())
	return nil
}
