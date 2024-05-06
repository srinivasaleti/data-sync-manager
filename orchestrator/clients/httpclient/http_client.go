package httpClient

import "net/http"

// IClient represents an http client interface.
type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client implements http client interface methods.
type Client struct {
	IClient
}

// Do sends an HTTP request and returns an HTTP response.
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	httpClient := http.Client{}
	return httpClient.Do(req)
}
