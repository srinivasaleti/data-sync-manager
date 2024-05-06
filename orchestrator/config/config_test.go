package config

import (
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/syncmanager"
	"github.com/stretchr/testify/assert"
)

func TestConfi(t *testing.T) {
	t.Run("should handle error while reading file", func(t *testing.T) {
		config, err := GetConfig("random_file")
		assert.Error(t, err)
		assert.Nil(t, config)
	})

	t.Run("should read config", func(t *testing.T) {
		config, err := GetConfig("./mocks/config.yaml")
		assert.Nil(t, err)
		assert.Equal(t, config[0], syncmanager.SyncConfig{
			Cron: "* * * * * *",
			Source: connectors.Config{
				Type: "s3",
				Config: map[string]string{
					"accessKey": "access_key",
					"bucket":    "bucket",
					"region":    "region",
					"secretKey": "secret_key",
				},
			},
			Target: connectors.Config{
				Type: "filesystem",
				Config: map[string]string{
					"outdir": "./test",
				},
			},
		})
	})
}
