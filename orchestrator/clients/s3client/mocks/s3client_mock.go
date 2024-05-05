package mocks

import (
	"fmt"
	"net/http"

	"github.com/srinivasaleti/data-sync-manager/orchestrator/clients/s3client"
)

type MockS3Client struct {
	objectURL     string
	signErr       error
	signResponse  http.Header
	signedRequest *http.Request
	s3client.IClient
}

func (client *MockS3Client) Sign(req *http.Request) (http.Header, error) {
	client.signedRequest = req
	return client.signResponse, client.signErr
}

func (client *MockS3Client) SignErr(err error) {
	client.signErr = err
}

func (client *MockS3Client) GetObjectUrl(key string) string {
	return client.objectURL
}

func (client *MockS3Client) SetObjectUrl(url string) {
	client.objectURL = url
}

func (client MockS3Client) IsRequestSigned(url string) bool {
	fmt.Println(client.signedRequest.URL, client.signedRequest.URL.String())
	return client.signedRequest != nil && client.signedRequest.URL.String() == url
}

func (client *MockS3Client) Clear() {
	client.signErr = nil
	client.objectURL = ""
	client.signedRequest = nil
	client.signResponse = nil
}
