package api_test

import (
	"testing"

	"github.com/ghostsecurity/terraform-provider-ghost/internal/api"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	baseURL := "https://example.base.api"
	apiKey := "1234"

	client := api.NewClient(baseURL, apiKey)

	req, err := client.NewRequest("GET", "/v2/apis", nil)
	require.NoError(t, err)

	expectedHeader := "Bearer 1234"
	require.Equal(t, expectedHeader, req.Header.Get("Authorization"))

	require.Equal(t, "/v2/apis", req.URL.Path)
	require.Equal(t, "example.base.api", req.URL.Host)
	require.Equal(t, "https", req.URL.Scheme)
}
