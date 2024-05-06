package s3client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Config represents s3 configuration.
type Config struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
	Bucket    string `json:"bucket"`
}

// IClient represents an interface for interacting with S3.
type IClient interface {
	Sign(req *http.Request) (http.Header, error)
	GetObjectUrl(key string) string
	ListKeys(callback func(keys []string)) ([]string, error)
}

type S3Client struct {
	*Config
	IClient
}

// Sign signs an HTTP request for S3 authentication.
// It returns the signed request headers or an error if signing fails.
func (s3 *S3Client) Sign(req *http.Request) (http.Header, error) {
	if s3.Config == nil {
		return nil, errors.New("client config should not be empty")
	}
	creds := credentials.NewStaticCredentials(s3.AccessKey, s3.SecretKey, "")
	signer := v4.NewSigner(creds)
	return signer.Sign(req, nil, "s3", s3.Region, time.Now())
}

// GetObjectUrl returns the URL for accessing an object in the S3 bucket
// specified by the given key.
func (s3 *S3Client) GetObjectUrl(key string) string {
	return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", s3.Region, s3.Bucket, key)
}

// ListKeys returns a list of keys (object identifiers) in the S3 bucket.
// It recursively checks all the keys inside a folder and subfolder too.
// It returns the list of keys or an error if listing fails.
// It also accepts callback function which will be called after fetching all keys in a page.
func (client *S3Client) ListKeys(callback func(keys []string)) ([]string, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(client.Region),
		Credentials: credentials.NewStaticCredentials(
			client.AccessKey,
			client.SecretKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}
	svc := s3.New(session)
	var allKeys []string
	params := &s3.ListObjectsInput{
		Bucket:  aws.String(client.Bucket),
		MaxKeys: aws.Int64(2),
	}
	err = svc.ListObjectsPages(params, func(page *s3.ListObjectsOutput, lastPage bool) bool {
		keys := extractKeysFromPage(page)
		allKeys = append(allKeys, keys...)
		callback(keys)
		return !lastPage
	})
	if err != nil {
		return nil, err
	}
	return allKeys, nil
}

func extractKeysFromPage(page *s3.ListObjectsOutput) []string {
	var allKeys []string
	for _, obj := range page.Contents {
		objKey := *obj.Key
		if !strings.HasSuffix(objKey, "/") {
			allKeys = append(allKeys, *obj.Key)
		}
	}
	return allKeys
}

func parseS3ClientConfig(configInterface map[string]string) (*Config, error) {
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

func NewS3Client(configMap map[string]string) (*S3Client, error) {
	config, err := parseS3ClientConfig(configMap)
	if err != nil {
		return nil, err
	}
	return &S3Client{
		Config: config,
	}, nil
}
