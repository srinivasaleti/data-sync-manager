package syncmanager

import (
	"errors"
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/factory"
	connectorsmock "github.com/srinivasaleti/data-sync-manager/orchestrator/connectors/mocks"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	schedulermock "github.com/srinivasaleti/data-sync-manager/orchestrator/scheduler/mocks"
	"github.com/stretchr/testify/assert"
)

var mockFactory = factory.NewMockFactory()
var mockScheduler = schedulermock.NewMockScheduler()
var sourceConnector = &connectorsmock.MockConnector{}
var targetConnector = &connectorsmock.MockConnector{}
var syncManager = New(mockFactory, mockScheduler, logger.NewLogger())
var file = "some file"

func reset() {
	sourceConnector.Reset()
	targetConnector.Reset()
	mockScheduler.Reset()
}

func TestScheduleAJobToSyncData(t *testing.T) {
	t.Run("should handle connectors not found errors", func(t *testing.T) {
		reset()
		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{}), errConnectorsRequired)

		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{
			Source: "s3",
			Target: "local",
		}), factory.ErrConnectorNotFound)

		mockFactory.SetConnector("s3", &connectorsmock.MockConnector{})
		assert.Equal(t, syncManager.scheduleSyncData(SyncConfig{
			Source: "s3",
			Target: "local",
		}), factory.ErrConnectorNotFound)
	})

	t.Run("should schedule the job", func(t *testing.T) {
		mockFactory.SetConnector("s3", sourceConnector)
		mockFactory.SetConnector("local", targetConnector)
		sourceConnector.SetGetResponse(file)

		err := syncManager.scheduleSyncData(SyncConfig{
			Cron:      "* * * * * 1",
			Source:    "s3",
			Target:    "local",
			ObjectKey: "id",
		})

		assert.Nil(t, err)
		assert.Equal(t, mockScheduler.GetLatestCronExpression(), "* * * * * 1")

		mockScheduler.GetScheduledTask()()

		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.True(t, targetConnector.PutShouldBeCalledWith(file))
	})
}

func TestSyncData(t *testing.T) {
	t.Run("should throw error when there is an error while getting data from source connector", func(t *testing.T) {
		reset()
		sourceConnectorGetErr := errors.New("unable to get from source")
		sourceConnector.SetGetErr(sourceConnectorGetErr)

		err := syncManager.syncData(sourceConnector, targetConnector, "id")

		assert.Error(t, err)
		assert.Equal(t, err, sourceConnectorGetErr)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.Equal(t, targetConnector.NumberOfGetCalls(), 0)
	})

	t.Run("should throw error when there is an error while adding data to target connector", func(t *testing.T) {
		reset()
		targetConnectorPutErr := errors.New("unable to add to target")
		sourceConnector.SetGetResponse(file)
		targetConnector.SetPutErr(targetConnectorPutErr)

		err := syncManager.syncData(sourceConnector, targetConnector, "id")

		assert.Error(t, err)
		assert.Equal(t, err, targetConnectorPutErr)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.Equal(t, sourceConnector.NumberOfGetCalls(), 1)
		assert.Equal(t, targetConnector.NumberOfPutCalls(), 1)
	})

	t.Run("should sync data", func(t *testing.T) {
		reset()
		sourceConnector.SetGetResponse(file)

		err := syncManager.syncData(sourceConnector, targetConnector, "id")

		assert.Nil(t, err)
		assert.True(t, sourceConnector.GetShouldBeCalledWith("id"))
		assert.True(t, targetConnector.PutShouldBeCalledWith(file))
		assert.Equal(t, sourceConnector.NumberOfGetCalls(), 1)
		assert.Equal(t, targetConnector.NumberOfPutCalls(), 1)
	})
}
