package s3

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	httpClient "github.com/srinivasaleti/data-sync-manager/orchestrator/clients/httpclient"
	s3clientmocks "github.com/srinivasaleti/data-sync-manager/orchestrator/clients/s3client/mocks"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
	"github.com/stretchr/testify/assert"
)

var mockLogger = logger.NewLogger()
var mockS3Client = &s3clientmocks.MockS3Client{}
var mockHttpClient = &httpClient.MockHttpClient{}
var connector, _ = New(mockLogger, mockS3Client, mockHttpClient)

func reset() {
	mockHttpClient.Clear()
	mockS3Client.Clear()
}

func TestToString(t *testing.T) {
	assert.Equal(t, connector.ToString(), "s3")
}

func TestGet(t *testing.T) {
	key := "test-key"
	objectUrl := "http://s3.objurl.com"
	responseBody := "some response body"

	t.Run("should make get request", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		mockHttpClient.SetResponse(&http.Response{
			Body: io.NopCloser(bytes.NewReader([]byte(responseBody))),
		})

		connector.Get(key)

		assert.Equal(t, mockHttpClient.GetRequest().URL.String(), objectUrl)
	})

	t.Run("should sign the request", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		mockHttpClient.SetResponse(&http.Response{
			Body: io.NopCloser(bytes.NewReader([]byte(responseBody))),
		})

		connector.Get(key)

		assert.True(t, mockS3Client.IsRequestSigned(objectUrl))
	})

	t.Run("should return error when signing request fails", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		mockHttpClient.SetResponse(&http.Response{
			Body: io.NopCloser(bytes.NewReader([]byte(responseBody))),
		})
		s3SignErr := errors.New("error during signing object")
		mockS3Client.SignErr(s3SignErr)

		data, err := connector.Get(key)

		assert.Nil(t, data)
		assert.Equal(t, err, s3SignErr)
	})

	t.Run("should return error when s3 get request fails", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		httpErr := errors.New("http err")
		mockHttpClient.SetErr(httpErr)

		data, err := connector.Get(key)

		assert.Nil(t, data)
		assert.Equal(t, err, httpErr)
	})

	t.Run("should return error when s3 return error status code", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		s3Err := s3.Error{Message: aws.String("not found")}
		s3ErrBytes, _ := json.Marshal(s3Err)
		mockHttpClient.SetResponse(&http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(bytes.NewReader([]byte(s3ErrBytes))),
		})

		data, err := connector.Get(key)

		assert.Nil(t, data)
		assert.Error(t, err)
	})

	t.Run("should return the request body", func(t *testing.T) {
		reset()
		mockS3Client.SetObjectUrl(objectUrl)
		mockHttpClient.SetResponse(&http.Response{
			Body: io.NopCloser(bytes.NewReader([]byte(responseBody))),
		})

		data, err := connector.Get(key)

		assert.NoError(t, err)
		assert.Equal(t, data, []byte(responseBody))
	})
}

func TestParseS3Err(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError s3.Error
	}{
		{
			name: "Valid S3 Error",
			input: `
				<Error>
					<Code>NoSuchKey</Code>
					<Message>The specified key does not exist.</Message>
				</Error>`,
			expectedError: s3.Error{
				Code:    aws.String("NoSuchKey"),
				Message: aws.String("The specified key does not exist."),
			},
		},
		{
			name: "Empty Error Message",
			input: `
				<Error>
					<Code>UnknownError</Code>
				</Error>`,
			expectedError: s3.Error{
				Message: aws.String("unknown error"),
			},
		},
		{
			name:          "Invalid XML",
			input:         "invalid-xml",
			expectedError: s3.Error{Message: aws.String("unable to parse s3 error")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			body := io.NopCloser(reader)

			result := parseS3Err(body)

			assert.Equal(t, tc.expectedError, result)
		})
	}
}
