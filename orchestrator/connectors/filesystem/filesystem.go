package filesystem

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
)

type Config struct {
	OutDirectory string `json:"outdir"`
}

type FileSystemConnector struct {
	connectors.Connector
	logger logger.ILogger
	Config
}

func (connector *FileSystemConnector) Get(key string) ([]byte, error) {
	return nil, errors.New("filesystem get not implemented yet")
}

func (connector *FileSystemConnector) Put(key string, data []byte) error {
	keyPath := filepath.Join(connector.OutDirectory, key)
	if err := os.MkdirAll(filepath.Dir(keyPath), 0755); err != nil {
		connector.logger.Error(err, "error while creating directory")
		return nil
	}
	// Create the output file
	outFile, err := os.Create(keyPath)
	if err != nil {
		connector.logger.Error(err, "error creating output file")
		return nil
	}
	defer outFile.Close()

	// Write to file
	_, err = outFile.Write(data)
	if err != nil {
		connector.logger.Error(err, "error writing payload to file")
		return nil
	}
	return nil
}

func parseConfig(configMap map[string]string) (*Config, error) {
	if configMap == nil {
		return nil, errors.New("configuration should not be nil")
	}
	configBytes, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err = json.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}
	return &config, err
}

func (connector *FileSystemConnector) ToString() string {
	return "filesystem"
}

func (connector *FileSystemConnector) Exists(key string) bool {
	keyPath := filepath.Join(connector.OutDirectory, key)
	_, err := os.Stat(keyPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func New(logger logger.ILogger, configMap map[string]string) (*FileSystemConnector, error) {
	config, err := parseConfig(configMap)
	if err != nil {
		return nil, err
	}
	return &FileSystemConnector{
		logger: logger,
		Config: *config,
	}, nil
}
