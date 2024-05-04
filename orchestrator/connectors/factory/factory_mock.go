package factory

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

type MockFactory struct {
	getConnectorErr error
	connectors      map[string]connectors.Connector
	IFactory
}

func (f *MockFactory) GetConnector(connector string) (connectors.Connector, error) {
	if f.connectors[connector] == nil {
		return nil, ErrConnectorNotFound
	}
	return f.connectors[connector], nil
}

func (f *MockFactory) SetGetConnectorErr(err error) {
	f.getConnectorErr = err
}

func (f *MockFactory) SetGetConnector(connectorString string, connector connectors.Connector) {
	f.connectors[connectorString] = connector
}

func NewMockFactory() *MockFactory {
	return &MockFactory{
		connectors: map[string]connectors.Connector{},
	}
}
