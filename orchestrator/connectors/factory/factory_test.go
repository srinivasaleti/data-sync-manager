package factory

import (
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	factory := New(logger.NewLogger())

	t.Run("return s3", func(t *testing.T) {
		connector, err := factory.GetConnector(connectors.Config{
			Type:   "s3",
			Config: map[string]string{},
		})
		assert.NoError(t, err)
		assert.Equal(t, "s3", connector.ToString())
	})

	t.Run("return error when  s3 is invalid", func(t *testing.T) {
		_, err := factory.GetConnector(connectors.Config{
			Type:   "s3",
			Config: nil,
		})
		assert.Error(t, err)
	})

	t.Run("return filesystem", func(t *testing.T) {
		connector, err := factory.GetConnector(connectors.Config{
			Type:   "filesystem",
			Config: map[string]string{},
		})
		assert.NoError(t, err)
		assert.Equal(t, "filesystem", connector.ToString())
	})

	t.Run("return connector not found", func(t *testing.T) {
		_, err := factory.GetConnector(connectors.Config{
			Type:   "unknow",
			Config: map[string]string{},
		})
		assert.Error(t, err)
		assert.Equal(t, err, ErrConnectorNotFound)
	})
}
