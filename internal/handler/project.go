package handler

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/cms"
	"github.com/reearth/reearth-cms-mcp/internal/tool"
	"github.com/reearth/reearthx/asset/domain/project"
)

func (h *Handler) GetProject(ctx context.Context, req *mcp.CallToolRequest, args GetProjectArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	resp, err := h.client.ProjectGetWithResponse(ctx, h.workspaceID, projectID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get project: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Project not found: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) ListProjects(ctx context.Context, req *mcp.CallToolRequest, args ListProjectsArgs) (*mcp.CallToolResult, any, error) {
	params := &cms.ProjectFilterParams{}

	if args.Page > 0 {
		page := cms.PageParam(args.Page)
		params.Page = &page
	}
	if args.PerPage > 0 {
		perPage := cms.PerPageParam(args.PerPage)
		params.PerPage = &perPage
	}
	if args.Keyword != "" {
		keyword := cms.KeywordParam(args.Keyword)
		params.Keyword = &keyword
	}

	resp, err := h.client.ProjectFilterWithResponse(ctx, h.workspaceID, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list projects: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list projects: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) UpdateProject(ctx context.Context, req *mcp.CallToolRequest, args UpdateProjectArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	body := cms.ProjectUpdateJSONRequestBody{}

	if args.Name != "" {
		body.Name = &args.Name
	}
	if args.Description != "" {
		body.Description = &args.Description
	}
	if args.Alias != "" {
		body.Alias = &args.Alias
	}

	resp, err := h.client.ProjectUpdateWithResponse(ctx, h.workspaceID, projectID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update project: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update project: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}
