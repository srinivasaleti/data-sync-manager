package s3

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	httpClient "github.com/srinivasaleti/data-sync-manager/orchestrator/clients/httpclient"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/clients/s3client"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/connectors"
	"github.com/srinivasaleti/data-sync-manager/orchestrator/logger"
)

// Config represent S3 connector config

type S3Connector struct {
	connectors.Connector
	S3Client   s3client.IClient
	HttpClient httpClient.IClient
	Logger     logger.ILogger
}

func (connector *S3Connector) Get(key string) ([]byte, error) {
	// Create the URL for the S3 object
	req, err := http.NewRequest("GET", connector.S3Client.GetObjectUrl(key), nil)
	if err != nil {
		connector.Logger.Error(err, "unable to create request", "id", key)
		return nil, err
	}
	_, err = connector.S3Client.Sign(req)
	if err != nil {
		connector.Logger.Error(err, "unable to sign request", "id", key)
		return nil, err
	}
	// Send the request
	response, err := connector.HttpClient.Do(req)
	if err != nil {
		connector.Logger.Error(err, "unable to send request", "id", key)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 && response.StatusCode <= 599 {
		s3Err := parseS3Err(response.Body)
		err := errors.New(*s3Err.Message)
		connector.Logger.Error(err, "error response from s3", "id", key, "status", response.StatusCode, "code", s3Err.Code, "message", s3Err.Message)
		return nil, err
	}
	return io.ReadAll(response.Body)
}

func (connector *S3Connector) Put(key string, data []byte) error {
	return nil
}

func (connector *S3Connector) ToString() string {
	return "s3"
}

func parseS3Err(body io.ReadCloser) s3.Error {
	var s3Err s3.Error
	if err := xml.NewDecoder(body).Decode(&s3Err); err != nil {
		return s3.Error{Message: aws.String("unable to parse s3 error")}
	}
	if s3Err.Message != nil {
		return s3Err
	}
	return s3.Error{Message: aws.String("unknown error")}
}

func New(logger logger.ILogger, client s3client.IClient, httpClient httpClient.IClient) (*S3Connector, error) {
	return &S3Connector{
		HttpClient: httpClient,
		Logger:     logger,
		S3Client:   client,
	}, nil
}
