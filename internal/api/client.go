package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("resource not found")

type GhostClient struct {
	BaseURL string
	APIKey  string
	HTTP    *http.Client
}

func NewClient(baseURL, apiKey string) *GhostClient {
	return &GhostClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTP: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// NewRequest returns a new HTTP request using the configured BaseURL and API key.
// The path provided should be relative to the base URL.
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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

type InvalidInputError struct {
	Err string `json:"error"`
}

type LogForwarderRequest struct {
	Name string `json:"name"`
}

type LogForwarder struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	SubjectID string           `json:"subject_id"`
	Source    *ForwarderSource `json:"source,omitempty"`
}

type ForwarderSource struct {
	BucketName string `json:"bucket_name"`
	RoleARN    string `json:"role_arn"`
	SQSARN     string `json:"sqs_arn"`
	Region     string `json:"region"`
	AccountID  string `json:"account_id"`
}

// CreateLogForwarder makes an authenticated request to the Ghost API to create a new log
// forwarder.
func (c *GhostClient) CreateLogForwarder(req LogForwarderRequest) (*LogForwarder, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("create log forwarder: %w", err)
	}

	body := bytes.NewReader(bodyBytes)
	createReq, err := c.NewRequest("POST", "/v2/log_forwarders", body)
	if err != nil {
		return nil, fmt.Errorf("create log forwarder: %w", err)
	}

	resp, err := c.HTTP.Do(createReq)
	if err != nil {
		return nil, fmt.Errorf("create log forwarder: %w", err)
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusCreated:
		var forwarder LogForwarder
		if err := dec.Decode(&forwarder); err != nil {
			return nil, fmt.Errorf("create log forwarder: %w", err)
		}
		return &forwarder, nil
	case http.StatusBadRequest:
		var inputErr InvalidInputError
		if err := dec.Decode(&inputErr); err != nil {
			return nil, fmt.Errorf("create log forwarder: %w", err)
		}
		return nil, fmt.Errorf("create log forwarder: %v", inputErr.Err)
	default:
		return nil, fmt.Errorf("create log forwarder: unexpected status %v", resp.StatusCode)
	}
}

// GetLogForwarder makes an authenticated request to the Ghost API to fetch the
// log forwarder with the given ID.
func (c *GhostClient) GetLogForwarder(id uuid.UUID) (*LogForwarder, error) {
	getReq, err := c.NewRequest("GET", "/v2/log_forwarders/"+id.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("get log forwarder: %w", err)
	}

	resp, err := c.HTTP.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("get log forwarder: %w", err)
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		var forwarder LogForwarder
		if err := dec.Decode(&forwarder); err != nil {
			return nil, fmt.Errorf("get log forwarder: %w", err)
		}
		return &forwarder, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("get log forwarder: %w", ErrNotFound)
	default:
		return nil, fmt.Errorf("get log forwarder: unexpected status %v", resp.StatusCode)
	}
}

// DeleteLogForwarder makes an authenticated request to the Ghost API to delete the
// log forwarder with the given ID.
func (c *GhostClient) DeleteLogForwarder(id uuid.UUID) error {
	deleteReq, err := c.NewRequest("DELETE", "/v2/log_forwarders/"+id.String(), nil)
	if err != nil {
		return fmt.Errorf("delete log forwarder: %w", err)
	}

	resp, err := c.HTTP.Do(deleteReq)
	if err != nil {
		return fmt.Errorf("delete log forwarder: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("delete log forwarder: %w", ErrNotFound)
	default:
		return fmt.Errorf("delete log forwarder: unexpected status %v", resp.StatusCode)
	}
}

type updateForwarderRequest struct {
	Source *ForwarderSource `json:"source,omitempty"`
}

// UpdateLogForwarderSource makes an authenticated request to the Ghost API to update the
// source for log forwarder with the given ID.
// Setting the source to nil deletes the source.
func (c *GhostClient) UpdateLogForwarderSource(id uuid.UUID, source *ForwarderSource) (*LogForwarder, error) {
	bodyBytes, err := json.Marshal(updateForwarderRequest{Source: source})
	if err != nil {
		return nil, fmt.Errorf("create log forwarder: %w", err)
	}

	body := bytes.NewReader(bodyBytes)
	patchReq, err := c.NewRequest("PATCH", "/v2/log_forwarders/"+id.String(), body)
	if err != nil {
		return nil, fmt.Errorf("update log forwarder source: %w", err)
	}

	resp, err := c.HTTP.Do(patchReq)
	if err != nil {
		return nil, fmt.Errorf("update log forwarder source: %w", err)
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		var forwarder LogForwarder
		if err := dec.Decode(&forwarder); err != nil {
			return nil, fmt.Errorf("update log forwarder: %w", err)
		}
		return &forwarder, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("update log forwarder: %w", ErrNotFound)
	default:
		bodyBytes, _ = io.ReadAll(resp.Body)
		return nil, fmt.Errorf("update log forwarder: unexpected status %v: %v", resp.StatusCode, string(bodyBytes))
	}
}
