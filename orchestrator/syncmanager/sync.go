package syncmanager

import (
	"errors"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

type SyncConfig struct {
	Source string
	Target string
}

func (s *SyncManager) scheduleSyncData(config SyncConfig) error {
	if len(config.Source) == 0 || len(config.Target) == 0 {
		return errors.New("source and target connectors are required")
	}
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
	s.syncData(sourceConnector, targetConnector)
	return nil
}

func (s *SyncManager) syncData(source connectors.Connector, target connectors.Connector) error {
	s.logger.Info("syncing data", "source", source.ToString(), "target", target.ToString())
	data, err := source.Get("id")
	s.logger.Info("successfully got data from source")
	if err != nil {
		s.logger.Error(err, "unable to get the data from source")
		return err
	}
	s.logger.Info("adding data to target")
	err = target.Put(data)
	if err != nil {
		s.logger.Error(err, "unable to create the data in target")
		return err
	}
	s.logger.Info("successfully synced data", "source", source.ToString(), "target", target.ToString())
	return nil
}