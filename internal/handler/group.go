package handler

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/cms"
	"github.com/reearth/reearth-cms-mcp/internal/tool"
	"github.com/reearth/reearthx/asset/domain/group"
	"github.com/reearth/reearthx/asset/domain/project"
)

func (h *Handler) GetGroup(ctx context.Context, req *mcp.CallToolRequest, args GetGroupArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	groupIDOrKey := group.IDOrKey(args.GroupID)

	resp, err := h.client.GroupGetWithResponse(ctx, h.workspaceID, projectID, groupIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get group: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Group not found: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) ListGroups(ctx context.Context, req *mcp.CallToolRequest, args ListGroupsArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	params := &cms.GroupFilterParams{}

	if args.Page > 0 {
		page := cms.PageParam(args.Page)
		params.Page = &page
	}
	if args.PerPage > 0 {
		perPage := cms.PerPageParam(args.PerPage)
		params.PerPage = &perPage
	}

	resp, err := h.client.GroupFilterWithResponse(ctx, h.workspaceID, projectID, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list groups: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list groups: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) CreateGroup(ctx context.Context, req *mcp.CallToolRequest, args CreateGroupArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	body := cms.GroupCreateJSONRequestBody{
		Name: args.Name,
		Key:  args.Key,
	}

	if args.Description != "" {
		body.Description = &args.Description
	}

	resp, err := h.client.GroupCreateWithResponse(ctx, h.workspaceID, projectID, nil, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create group: %v", err)), nil, nil
	}
	if resp.JSON201 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create group: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON201), nil, nil
}

func (h *Handler) UpdateGroup(ctx context.Context, req *mcp.CallToolRequest, args UpdateGroupArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	groupIDOrKey := group.IDOrKey(args.GroupID)

	body := cms.GroupUpdateJSONRequestBody{}

	if args.Name != "" {
		body.Name = &args.Name
	}
	if args.Key != "" {
		body.Key = &args.Key
	}
	if args.Description != "" {
		body.Description = &args.Description
	}

	resp, err := h.client.GroupUpdateWithResponse(ctx, h.workspaceID, projectID, groupIDOrKey, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update group: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update group: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteGroup(ctx context.Context, req *mcp.CallToolRequest, args DeleteGroupArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	groupIDOrKey := group.IDOrKey(args.GroupID)

	resp, err := h.client.GroupDeleteWithResponse(ctx, h.workspaceID, projectID, groupIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete group: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete group: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Group deleted successfully"), nil, nil
}
