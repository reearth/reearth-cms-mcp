package handler

import (
	"context"
	"fmt"

	"github.com/KeisukeYamashita/reearth-cms-mcp/internal/cms"
	"github.com/KeisukeYamashita/reearth-cms-mcp/internal/tool"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearthx/asset/domain/model"
	"github.com/reearth/reearthx/asset/domain/project"
)

func (h *Handler) GetModel(ctx context.Context, req *mcp.CallToolRequest, args GetModelArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	resp, err := h.client.ModelGetWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get model: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Model not found: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) ListModels(ctx context.Context, req *mcp.CallToolRequest, args ListModelsArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	params := &cms.ModelFilterParams{}

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

	resp, err := h.client.ModelFilterWithResponse(ctx, h.workspaceID, projectID, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list models: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list models: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) CreateModel(ctx context.Context, req *mcp.CallToolRequest, args CreateModelArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	body := cms.ModelCreateJSONRequestBody{
		Name: &args.Name,
	}

	if args.Key != "" {
		body.Key = &args.Key
	}
	if args.Description != "" {
		body.Description = &args.Description
	}

	resp, err := h.client.ModelCreateWithResponse(ctx, h.workspaceID, projectID, nil, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create model: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create model: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) UpdateModel(ctx context.Context, req *mcp.CallToolRequest, args UpdateModelArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	body := cms.ModelUpdateJSONRequestBody{}

	if args.Name != "" {
		body.Name = &args.Name
	}
	if args.Key != "" {
		body.Key = &args.Key
	}
	if args.Description != "" {
		body.Description = &args.Description
	}

	resp, err := h.client.ModelUpdateWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update model: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update model: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteModel(ctx context.Context, req *mcp.CallToolRequest, args DeleteModelArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	resp, err := h.client.ModelDeleteWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete model: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete model: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Model deleted successfully"), nil, nil
}

func (h *Handler) CopyModel(ctx context.Context, req *mcp.CallToolRequest, args CopyModelArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	body := cms.ModelCopyJSONRequestBody{}

	if args.Name != "" {
		body.Name = &args.Name
	}
	if args.Key != "" {
		body.Key = &args.Key
	}

	resp, err := h.client.ModelCopyWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to copy model: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to copy model: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) GetModelSchema(ctx context.Context, req *mcp.CallToolRequest, args GetModelSchemaArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	resp, err := h.client.SchemaByModelAsJSONWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get model schema: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get model schema: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) GetModelMetadataSchema(ctx context.Context, req *mcp.CallToolRequest, args GetModelMetadataSchemaArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	resp, err := h.client.MetadataSchemaByModelAsJSONWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get metadata schema: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get metadata schema: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}
