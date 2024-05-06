package s3client

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestParseS3ClientConfig(t *testing.T) {
	testCases := []struct {
		name        string
		configMap   map[string]string
		expected    *Config
		expectedErr error
	}{
		{
			name: "ValidConfig",
			configMap: map[string]string{
				"accessKey": "access_key",
				"secretKey": "secret_key",
				"region":    "region",
				"bucket":    "bucket",
			},
			expected: &Config{
				AccessKey: "access_key",
				SecretKey: "secret_key",
				Region:    "region",
				Bucket:    "bucket",
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
			config, err := parseS3ClientConfig(tc.configMap)
			assert.Equal(t, err, tc.expectedErr)
			assert.Equal(t, config, tc.expected)
		})
	}
}

func TestExtractKeysFromPage(t *testing.T) {
	tests := []struct {
		name     string
		page     *s3.ListObjectsOutput
		expected []string
	}{
		{
			name:     "NoObjects",
			page:     &s3.ListObjectsOutput{},
			expected: nil,
		},
		{
			name: "ObjectsWithSuffix",
			page: &s3.ListObjectsOutput{
				Contents: []*s3.Object{
					{Key: aws.String("file1.txt")},
					{Key: aws.String("file2.txt")},
				},
			},
			expected: []string{"file1.txt", "file2.txt"},
		},
		{
			name: "ObjectsWithoutSuffix",
			page: &s3.ListObjectsOutput{
				Contents: []*s3.Object{
					{Key: aws.String("folder1/")},
					{Key: aws.String("folder2/file3.txt")},
				},
			},
			expected: []string{"folder2/file3.txt"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := extractKeysFromPage(test.page)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestGetObjectUrl(t *testing.T) {
	client, err := NewS3Client(map[string]string{
		"accessKey": "access_key",
		"secretKey": "secret_key",
		"region":    "region",
		"bucket":    "bucket",
	})
	assert.NoError(t, err)
	assert.Equal(t, client.GetObjectUrl("1"), "https://s3.region.amazonaws.com/bucket/1")
	assert.Equal(t, client.GetObjectUrl("2"), "https://s3.region.amazonaws.com/bucket/2")
}
