package factory

import (
	httpClient "github.com/srinivasaleti/data-sync-manager/orchestrator/clients/httpclient"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/clients/s3client"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/filesystem"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/s3"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
)

type Factory struct {
	IFactory
	logger logger.ILogger
}

type IFactory interface {
	GetConnector(connector connectors.Config) (connectors.Connector, error)
}

func (f *Factory) GetConnector(connector connectors.Config) (connectors.Connector, error) {
	if connector.Type == "s3" {
		s3Client, err := s3client.NewS3Client(connector.Config)
		if err != nil {
			return nil, err
		}
		return s3.New(f.logger, s3Client, &httpClient.Client{})
	}
	if connector.Type == "filesystem" {
		return filesystem.New(f.logger, connector.Config)
	}
	return nil, ErrConnectorNotFound
}

func New(logger logger.ILogger) *Factory {
	return &Factory{
		logger: logger,
	}
}
