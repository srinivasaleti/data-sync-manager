package httpClient

import (
	"net/http"

	httpClient "github.com/srinivasaleti/data-sync-manager/orchestrator/clients/httpclient"
)

// MockHttpClient mocks http client.
type MockHttpClient struct {
	response *http.Response
	request  *http.Request
	err      error
	httpClient.IClient
}

func (client *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	client.request = req
	return client.response, client.err
}

func (client *MockHttpClient) SetErr(err error) {
	client.err = err
}

func (client *MockHttpClient) SetResponse(response *http.Response) {
	client.response = response
}
func (client *MockHttpClient) GetRequest() *http.Request {
	return client.request
}

func (client *MockHttpClient) Clear() {
	client.response = nil
	client.err = nil
}
