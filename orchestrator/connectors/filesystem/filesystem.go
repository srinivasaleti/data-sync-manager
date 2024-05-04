package filesystem

import (
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
)

type FileSystemConnector struct {
	connectors.Connector
}

func (s *FileSystemConnector) Get(id string) (interface{}, error) {
	return "hello", nil
}

func (s *FileSystemConnector) Put(data interface{}) error {
	return nil
}

func (s *FileSystemConnector) ToString() string {
	return "local"
}

func New() *FileSystemConnector {
	return &FileSystemConnector{}
}
