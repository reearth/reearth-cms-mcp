package handler

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/cms"
	"github.com/reearth/reearth-cms-mcp/internal/tool"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/asset/domain/project"
	"github.com/reearth/reearthx/asset/domain/schema"
)

func (h *Handler) CreateField(ctx context.Context, req *mcp.CallToolRequest, args CreateFieldArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	schemaID, err := id.SchemaIDFrom(args.SchemaID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid schema ID: %v", err)), nil, nil
	}

	valueType := cms.ValueType(args.Type)
	body := cms.FieldCreateJSONRequestBody{
		Key:      &args.Key,
		Type:     &valueType,
		Multiple: &args.Multiple,
		Required: &args.Required,
	}

	resp, err := h.client.FieldCreateWithResponse(ctx, h.workspaceID, projectID, schemaID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create field: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create field: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) UpdateField(ctx context.Context, req *mcp.CallToolRequest, args UpdateFieldArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	schemaID, err := id.SchemaIDFrom(args.SchemaID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid schema ID: %v", err)), nil, nil
	}

	fieldIDOrKey := schema.FieldIDOrKey(args.FieldID)

	body := cms.FieldUpdateJSONRequestBody{}

	if args.Key != "" {
		body.Key = &args.Key
	}
	if args.Type != "" {
		valueType := cms.ValueType(args.Type)
		body.Type = &valueType
	}
	if args.Multiple != nil {
		body.Multiple = args.Multiple
	}
	if args.Required != nil {
		body.Required = args.Required
	}

	resp, err := h.client.FieldUpdateWithResponse(ctx, h.workspaceID, projectID, schemaID, fieldIDOrKey, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update field: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update field: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteField(ctx context.Context, req *mcp.CallToolRequest, args DeleteFieldArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)

	schemaID, err := id.SchemaIDFrom(args.SchemaID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid schema ID: %v", err)), nil, nil
	}

	fieldIDOrKey := schema.FieldIDOrKey(args.FieldID)

	resp, err := h.client.FieldDeleteWithResponse(ctx, h.workspaceID, projectID, schemaID, fieldIDOrKey)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete field: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete field: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Field deleted successfully"), nil, nil
}
