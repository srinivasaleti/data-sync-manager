package httpClient

import "net/http"

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	IClient
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	httpClient := http.Client{}
	return httpClient.Do(req)
}
