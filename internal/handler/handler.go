package handler

import "github.com/reearth/reearthx/account/accountdomain"

type Handler struct {
	client      *CMSClient
	workspaceID accountdomain.WorkspaceIDOrAlias
}

func New(cfg Config) (*Handler, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Handler{
		client:      client,
		workspaceID: accountdomain.WorkspaceIDOrAlias(cfg.WorkspaceID),
	}, nil
}
