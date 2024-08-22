package http_client

import (
	"errors"
	"net/http"
	"time"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

type client struct {
	baseUrl    string
	httpClient *http.Client
}

// Creates a new Client instance
func NewClient(baseUrl string) (*client, error) {
	if baseUrl == "" {
		return nil, errors.New("Invalid baseUrl provided")
	}

	httpClient := http.Client{Timeout: time.Duration(5) * time.Second}
	return &client{
		baseUrl:    baseUrl,
		httpClient: &httpClient,
	}, nil
}
