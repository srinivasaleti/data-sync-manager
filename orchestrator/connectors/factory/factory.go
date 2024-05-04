package factory

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/filesystem"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/s3"
)

type Factory struct {
	IFactory
}

type IFactory interface {
	GetConnector(connector string) (connectors.Connector, error)
}

func (f *Factory) GetConnector(connector string) (connectors.Connector, error) {
	if connector == "s3" {
		return s3.New(), nil
	}
	if connector == "local" {
		return filesystem.New(), nil
	}
	return nil, ErrConnectorNotFound
}

func New() *Factory {
	return &Factory{}
}
