package config

import (
	"errors"
	"os"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"
	"gopkg.in/yaml.v2"
)

type Config struct{}

var (
	errInvalidSourceType              = errors.New("source type must be 's3'")
	errInvalidTargetType              = errors.New("target type must be 'filesystem'")
	errCronExpressionShouldNotBeEmpty = errors.New("cron expression must be provided")
)

func GetConfig(filePath string) (*syncmanager.SyncConfig, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var syncConfig syncmanager.SyncConfig
	err = yaml.Unmarshal(yamlFile, &syncConfig)
	if err != nil {
		return nil, err
	}
	if err := validate(&syncConfig); err != nil {
		return nil, err
	}
	return &syncConfig, nil
}
func validate(config *syncmanager.SyncConfig) error {
	if len(config.Cron) == 0 {
		return errCronExpressionShouldNotBeEmpty
	}
	if config.Source.Type != "s3" {
		return errInvalidSourceType
	}
	if config.Target.Type != "filesystem" {
		return errInvalidTargetType
	}
	return nil
}
