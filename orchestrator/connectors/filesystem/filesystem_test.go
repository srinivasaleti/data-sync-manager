package filesystem

import (
	"errors"
	"os"
	"testing"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	connector, err := New(logger.NewLogger(), map[string]string{
		"outdir": "",
	})
	assert.NoError(t, err)
	key := "test.txt"
	data := []byte("This is a test data.")

	err = connector.Put(key, data)
	assert.NoError(t, err)

	fileContent, err := os.ReadFile(key)
	assert.NoError(t, err)
	assert.Equal(t, data, fileContent)
	os.Remove(key)
}

func TestParseConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configMap   map[string]string
		expected    *Config
		expectedErr error
	}{
		{
			name: "ValidConfig",
			configMap: map[string]string{
				"outdir": "/path/to/output",
			},
			expected: &Config{
				OutDirectory: "/path/to/output",
			},
			expectedErr: nil,
		},
		{
			name:        "NilConfig",
			configMap:   nil,
			expected:    nil,
			expectedErr: errors.New("configuration should not be nil"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config, err := parseConfig(tc.configMap)
			assert.Equal(t, err, tc.expectedErr)
			assert.Equal(t, config, tc.expected)
		})
	}
}
