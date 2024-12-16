package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GhostClient struct {
	BaseURL string
	APIKey  string
}

func NewClient(baseURL, apiKey string) *GhostClient {
	return &GhostClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (c *GhostClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	requestURL, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	authHeader := "Bearer " + c.APIKey
	req.Header.Add("Authorization", authHeader)

	return req, nil
}
