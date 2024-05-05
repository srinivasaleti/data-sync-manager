package filesystem

import (
	"os"
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	connector := New(logger.NewLogger())
	key := "test.txt"
	data := []byte("This is a test data.")

	err := connector.Put(key, data)
	assert.NoError(t, err)

	fileContent, err := os.ReadFile(key)
	assert.NoError(t, err)
	assert.Equal(t, data, fileContent)
	os.Remove(key)
}
