package s3client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

type Config struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
	Bucket    string `json:"bucket"`
}

type IClient interface {
	Sign(req *http.Request) (http.Header, error)
	GetObjectUrl(key string) string
}

type S3Client struct {
	*Config
	IClient
}

func (s3 *S3Client) Sign(req *http.Request) (http.Header, error) {
	if s3.Config == nil {
		return nil, errors.New("client config should not be empty")
	}
	creds := credentials.NewStaticCredentials(s3.AccessKey, s3.SecretKey, "")
	signer := v4.NewSigner(creds)
	return signer.Sign(req, nil, "s3", s3.Region, time.Now())
}

func (s3 *S3Client) GetConfig() Config {
	return *s3.Config
}

func (s3 *S3Client) GetObjectUrl(key string) string {
	return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", s3.Region, s3.Bucket, key)
}

func parseS3ClientConfig(configInterface interface{}) (*Config, error) {
	if configInterface == nil {
		return nil, errors.New("configuration should not be nil")
	}
	configBytes, err := json.Marshal(configInterface)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err = json.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}
	return &config, err
}

func NewS3Client(configMap interface{}) (*S3Client, error) {
	config, err := parseS3ClientConfig(configMap)
	if err != nil {
		return nil, err
	}
	return &S3Client{
		Config: config,
	}, nil
}
