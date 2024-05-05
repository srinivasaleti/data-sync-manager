package filesystem

import (
	"os"
	"path/filepath"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
)

type FileSystemConnector struct {
	connectors.Connector
	logger logger.ILogger
}

func (connector *FileSystemConnector) Get(key string) ([]byte, error) {
	return nil, nil
}

func (s *FileSystemConnector) Put(key string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(key), 0755); err != nil {
		s.logger.Error(err, "error while creating directory")
		return nil
	}
	// Create the output file
	outFile, err := os.Create(key)
	if err != nil {
		s.logger.Error(err, "error creating output file")
		return nil
	}
	defer outFile.Close()

	// Write to file
	_, err = outFile.Write(data)
	if err != nil {
		s.logger.Error(err, "error writing payload to file")
		return nil
	}
	return nil
}

func (s *FileSystemConnector) ToString() string {
	return "local"
}

func New(logger logger.ILogger) *FileSystemConnector {
	return &FileSystemConnector{
		logger: logger,
	}
}
