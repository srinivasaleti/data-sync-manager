package factory

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
)

type MockFactory struct {
	getConnectorErr error
	connectors      map[string]connectors.Connector
	factory.IFactory
}

func (f *MockFactory) GetConnector(connector connectors.Config) (connectors.Connector, error) {
	if f.connectors[connector.Type] == nil {
		return nil, factory.ErrConnectorNotFound
	}
	return f.connectors[connector.Type], nil
}

func (f *MockFactory) SetGetConnectorErr(err error) {
	f.getConnectorErr = err
}

func (f *MockFactory) SetConnector(connectorString string, connector connectors.Connector) {
	f.connectors[connectorString] = connector
}

func NewMockFactory() *MockFactory {
	return &MockFactory{
		connectors: map[string]connectors.Connector{},
	}
}
