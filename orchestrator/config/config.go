package config

import (
	"os"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"
	"gopkg.in/yaml.v2"
)

type Config struct{}

func GetConfig(filePath string) (*syncmanager.SyncConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var contents syncmanager.SyncConfig
	err = yaml.Unmarshal(yamlFile, &contents)
	if err != nil {
		return nil, err
	}
	return &contents, nil
}
