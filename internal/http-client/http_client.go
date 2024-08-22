package http_client

import "net/http"

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

type client struct {
	baseUrl    string
	httpClient *http.Client
}
