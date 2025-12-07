package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KeisukeYamashita/reearth-cms-mcp/internal/cms"
)

type Config struct {
	Token       string
	BaseURL     string
	WorkspaceID string
}

type CMSClient struct {
	*cms.ClientWithResponses
}

func NewClient(cfg Config) (*CMSClient, error) {
	if cfg.Token == "" {
		return nil, fmt.Errorf("token is required")
	}
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("base_url is required")
	}

	authEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		return nil
	}

	client, err := cms.NewClientWithResponses(cfg.BaseURL, cms.WithRequestEditorFn(authEditor))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &CMSClient{
		ClientWithResponses: client,
	}, nil
}
