package handler

import (
	"context"
	"fmt"
	"io"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/cms"
	"github.com/reearth/reearth-cms-mcp/internal/logging"
	"github.com/reearth/reearth-cms-mcp/internal/tool"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/asset/domain/model"
	"github.com/reearth/reearthx/asset/domain/project"
)

func (h *Handler) GetItem(ctx context.Context, req *mcp.CallToolRequest, args GetItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	// assetParam := cms.AssetEmbedding(cms.All)
	// params := &cms.ItemGetParams{
	// 	Asset: (*cms.AssetParam)(&assetParam),
	// }

	resp, err := h.client.ItemGetWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, nil)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get item: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Item not found: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) ListItems(ctx context.Context, req *mcp.CallToolRequest, args ListItemsArgs) (*mcp.CallToolResult, any, error) {
	logger := logging.Logger(req.Session)
	logger.Info("listing items",
		"project_id", args.ProjectID,
		"model_id", args.ModelID,
		"page", args.Page,
		"per_page", args.PerPage,
		"keyword", args.Keyword,
	)

	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	assetParam := cms.AssetEmbedding(cms.All)
	params := &cms.ItemFilterParams{
		Asset: (*cms.AssetParam)(&assetParam),
	}

	if args.Page > 0 {
		pageParam := cms.PageParam(args.Page)
		params.Page = &pageParam
	}
	if args.PerPage > 0 {
		perPageParam := cms.PerPageParam(args.PerPage)
		params.PerPage = &perPageParam
	}
	if args.Keyword != "" {
		keyword := cms.KeywordParam(args.Keyword)
		params.Keyword = &keyword
	}

	logger.Debug("sending request to CMS API",
		"workspace_id", string(h.workspaceID),
		"project_id", string(projectID),
		"model_id", string(modelIDOrKey),
	)

	resp, err := h.client.ItemFilterWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, params)
	if err != nil {
		logger.Error("failed to list items", "error", err)
		return tool.ErrorResult(fmt.Sprintf("Failed to list items: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		logger.Error("failed to list items",
			"status_code", resp.StatusCode(),
			"body", string(resp.Body),
		)
		return tool.ErrorResult(fmt.Sprintf("Failed to list items: status %d", resp.StatusCode())), nil, nil
	}

	logger.Info("items listed successfully", "total_count", resp.JSON200.TotalCount)
	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) CreateItem(ctx context.Context, req *mcp.CallToolRequest, args CreateItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	fields := make([]cms.Field, len(args.Fields))
	for i, f := range args.Fields {
		key := f.Key
		valueType := cms.ValueType(f.Type)
		fields[i] = cms.Field{
			Key:   &key,
			Type:  &valueType,
			Value: f.Value,
		}
	}

	body := cms.ItemCreateJSONRequestBody{
		Fields: &fields,
	}

	resp, err := h.client.ItemCreateWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create item: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create item: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) UpdateItem(ctx context.Context, req *mcp.CallToolRequest, args UpdateItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	fields := make([]cms.Field, len(args.Fields))
	for i, f := range args.Fields {
		key := f.Key
		valueType := cms.ValueType(f.Type)
		fields[i] = cms.Field{
			Key:   &key,
			Type:  &valueType,
			Value: f.Value,
		}
	}

	body := cms.ItemUpdateJSONRequestBody{
		Fields: &fields,
	}

	resp, err := h.client.ItemUpdateWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update item: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update item: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteItem(ctx context.Context, req *mcp.CallToolRequest, args DeleteItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	resp, err := h.client.ItemDeleteWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete item: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete item: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Item deleted successfully"), nil, nil
}

func (h *Handler) CommentToItem(ctx context.Context, req *mcp.CallToolRequest, args CommentToItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	body := cms.ItemCommentCreateJSONRequestBody{
		Content: &args.Content,
	}

	resp, err := h.client.ItemCommentCreateWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to add comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to add comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment added successfully"), nil, nil
}

func (h *Handler) ListItemComments(ctx context.Context, req *mcp.CallToolRequest, args ListItemCommentsArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	resp, err := h.client.ItemCommentListWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list comments: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list comments: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteItemComment(ctx context.Context, req *mcp.CallToolRequest, args DeleteItemCommentArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	commentID, err := id.CommentIDFrom(args.CommentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid comment ID: %v", err)), nil, nil
	}

	resp, err := h.client.ItemCommentDeleteWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, commentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment deleted successfully"), nil, nil
}

func (h *Handler) UpdateItemComment(ctx context.Context, req *mcp.CallToolRequest, args UpdateItemCommentArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	commentID, err := id.CommentIDFrom(args.CommentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid comment ID: %v", err)), nil, nil
	}

	body := cms.ItemCommentUpdateJSONRequestBody{
		Content: &args.Content,
	}

	resp, err := h.client.ItemCommentUpdateWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, commentID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to update comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment updated successfully"), nil, nil
}

func (h *Handler) PublishItem(ctx context.Context, req *mcp.CallToolRequest, args PublishItemArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)
	itemID, err := id.ItemIDFrom(args.ItemID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid item ID: %v", err)), nil, nil
	}

	resp, err := h.client.ItemPublishWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, itemID, nil)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to publish item: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to publish item: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) GetItemsAsCSV(ctx context.Context, req *mcp.CallToolRequest, args GetItemsAsCSVArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	params := &cms.ItemsAsCSVParams{}
	if args.Page > 0 {
		page := cms.PageParam(args.Page)
		params.Page = &page
	}
	if args.PerPage > 0 {
		perPage := cms.PerPageParam(args.PerPage)
		params.PerPage = &perPage
	}

	resp, err := h.client.ItemsAsCSVWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get items as CSV: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to get items as CSV: status %d", resp.StatusCode())), nil, nil
	}

	body, err := io.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to read response body: %v", err)), nil, nil
	}

	return tool.TextResult(string(body)), nil, nil
}

func (h *Handler) GetItemsAsGeoJSON(ctx context.Context, req *mcp.CallToolRequest, args GetItemsAsGeoJSONArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	modelIDOrKey := model.IDOrKey(args.ModelID)

	params := &cms.ItemsAsGeoJSONParams{}
	if args.Page > 0 {
		page := cms.PageParam(args.Page)
		params.Page = &page
	}
	if args.PerPage > 0 {
		perPage := cms.PerPageParam(args.PerPage)
		params.PerPage = &perPage
	}

	resp, err := h.client.ItemsAsGeoJSONWithResponse(ctx, h.workspaceID, projectID, modelIDOrKey, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get items as GeoJSON: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to get items as GeoJSON: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult(string(resp.Body)), nil, nil
}
