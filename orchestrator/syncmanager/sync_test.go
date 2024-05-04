package syncmanager

import (
	"errors"
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/stretchr/testify/assert"
)

func TestScheduleAJobToSyncData(t *testing.T) {
	mockFactory := factory.NewMockFactory()
	syncManager := New(mockFactory, logger.NewLogger())

	t.Run("should handle connectors not found errors", func(t *testing.T) {
		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{}), errConnectorsRequired)

		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{
			Source: "s3",
			Target: "local",
		}), factory.ErrConnectorNotFound)

		mockFactory.SetGetConnector("s3", &connectors.MockConnector{})
		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{
			Source: "s3",
			Target: "local",
		}), factory.ErrConnectorNotFound)
	})
}

func TestSyncData(t *testing.T) {
	mockFactory := factory.NewMockFactory()
	syncManager := New(mockFactory, logger.NewLogger())
	sourceConnector := &connectors.MockConnector{}
	targetConnector := &connectors.MockConnector{}
	file := "somefile"

	t.Run("should throw error when there is an error while getting data from source connector", func(t *testing.T) {
		sourceConnector.Reset()
		targetConnector.Reset()
		sourceConnectorGetErr := errors.New("unable to get from source")
		sourceConnector.SetGetErr(sourceConnectorGetErr)

		err := syncManager.syncData(sourceConnector, targetConnector)

		assert.Error(t, err)
		assert.Equal(t, err, sourceConnectorGetErr)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.Equal(t, targetConnector.NumberOfGetCalls(), 0)
	})

	t.Run("should throw error when there is an error while adding data to target connector", func(t *testing.T) {
		sourceConnector.Reset()
		targetConnector.Reset()
		targetConnectorPutErr := errors.New("unable to add to target")
		sourceConnector.SetGetResponse(file)
		targetConnector.SetPutErr(targetConnectorPutErr)

		err := syncManager.syncData(sourceConnector, targetConnector)

		assert.Error(t, err)
		assert.Equal(t, err, targetConnectorPutErr)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.Equal(t, sourceConnector.NumberOfGetCalls(), 1)
		assert.Equal(t, targetConnector.NumberOfPutCalls(), 1)
	})

	t.Run("should sync data", func(t *testing.T) {
		sourceConnector.Reset()
		targetConnector.Reset()
		sourceConnector.SetGetResponse(file)

		err := syncManager.syncData(sourceConnector, targetConnector)

		assert.Nil(t, err)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.True(t, targetConnector.PutShouldBeCalledWith(file))
		assert.Equal(t, sourceConnector.NumberOfGetCalls(), 1)
		assert.Equal(t, targetConnector.NumberOfPutCalls(), 1)
	})
}
