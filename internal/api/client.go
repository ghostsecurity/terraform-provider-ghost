package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ghostsecurity/terraform-provider-ghost/internal/api/client"
)

var ErrNotFound = errors.New("resource not found")

type GhostClient struct {
	API *client.APIClient
}

func NewClient(baseURL, apiKey string) (*GhostClient, error) {
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("new client: %v", err)
	}

	// Configure the client to use the provided API key for all requests.
	cfg := client.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)

	// Set the base URL for the client
	cfg.Host = apiURL.Host
	cfg.Scheme = apiURL.Scheme
	cfg.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	apiClient := client.NewAPIClient(cfg)

	return &GhostClient{
		API: apiClient,
	}, nil
}

type LogForwarder struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SubjectID string `json:"subject_id"`
}

// CreateLogForwarder makes an authenticated request to the Ghost API to create a new log
// forwarder.
func (c *GhostClient) CreateLogForwarder(ctx context.Context, name string) (*LogForwarder, error) {
	req := c.API.LogForwardersAPI.
		CreateLogForwarder(ctx).
		LogForwarder(client.LogForwarderCreateParams{Name: name})

	forwarder, resp, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("create log forwarder: %w", err)
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		return &LogForwarder{
			ID:        forwarder.Id,
			Name:      forwarder.Name,
			SubjectID: forwarder.SubjectId,
		}, nil
	default:
		return nil, fmt.Errorf("create log forwarder: unexpected status %v", resp.StatusCode)
	}
}

// GetLogForwarder makes an authenticated request to the Ghost API to fetch the
// log forwarder with the given ID.
func (c *GhostClient) GetLogForwarder(ctx context.Context, id string) (*LogForwarder, error) {
	req := c.API.LogForwardersAPI.GetLogForwarder(ctx, id)

	forwarder, resp, err := req.Execute()
	if err != nil {
		// If the log forwarder with the given ID was not found return
		// a specific error type.
		var apiErr *client.GenericOpenAPIError
		if errors.As(err, &apiErr) {
			if _, ok := apiErr.Model().(client.NotFoundError); ok {
				return nil, fmt.Errorf("delete log forwarder: %w", ErrNotFound)
			}
		}
		return nil, fmt.Errorf("get log forwarder: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return &LogForwarder{
			ID:        forwarder.Id,
			Name:      forwarder.Name,
			SubjectID: forwarder.SubjectId,
		}, nil
	default:
		return nil, fmt.Errorf("get log forwarder: unexpected status %v", resp.StatusCode)
	}
}

// DeleteLogForwarder makes an authenticated request to the Ghost API to delete the
// log forwarder with the given ID.
func (c *GhostClient) DeleteLogForwarder(ctx context.Context, id string) error {
	req := c.API.LogForwardersAPI.DeleteLogForwarder(ctx, id)

	resp, err := req.Execute()
	if err != nil {
		// If the log forwarder with the given ID was not found return
		// a specific error type.
		var apiErr *client.GenericOpenAPIError
		if errors.As(err, &apiErr) {
			if _, ok := apiErr.Model().(client.NotFoundError); ok {
				return fmt.Errorf("delete log forwarder: %w", ErrNotFound)
			}
		}
		return fmt.Errorf("delete log forwarder: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("get log forwarder: unexpected status %v", resp.StatusCode)
	}
}
