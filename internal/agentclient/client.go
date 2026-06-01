package agentclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"openclaw/internal/domain"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) Run(ctx context.Context, req domain.RunRequest) (domain.RunResult, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return domain.RunResult{}, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/run", bytes.NewReader(body))
	if err != nil {
		return domain.RunResult{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return domain.RunResult{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return domain.RunResult{}, errors.New("agent service returned " + resp.Status)
	}
	var result domain.RunResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domain.RunResult{}, err
	}
	return result, nil
}
