package handler

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/cms"
	"github.com/reearth/reearth-cms-mcp/internal/tool"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/asset/domain/project"
)

func (h *Handler) GetAsset(ctx context.Context, req *mcp.CallToolRequest, args GetAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetGetWithResponse(ctx, h.workspaceID, projectID, assetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to get asset: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Asset not found: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) ListAssets(ctx context.Context, req *mcp.CallToolRequest, args ListAssetsArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	params := &cms.AssetFilterParams{}

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

	resp, err := h.client.AssetFilterWithResponse(ctx, h.workspaceID, projectID, params)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list assets: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list assets: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteAsset(ctx context.Context, req *mcp.CallToolRequest, args DeleteAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetDeleteWithResponse(ctx, h.workspaceID, projectID, assetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete asset: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete asset: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Asset deleted successfully"), nil, nil
}

func (h *Handler) CreateAsset(ctx context.Context, req *mcp.CallToolRequest, args CreateAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	body := cms.AssetCreateJSONRequestBody{
		Url: &args.URL,
	}

	resp, err := h.client.AssetCreateWithResponse(ctx, h.workspaceID, projectID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create asset: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to create asset: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) CommentToAsset(ctx context.Context, req *mcp.CallToolRequest, args CommentToAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	body := cms.AssetCommentCreateJSONRequestBody{
		Content: &args.Content,
	}

	resp, err := h.client.AssetCommentCreateWithResponse(ctx, h.workspaceID, projectID, assetID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to add comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to add comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment added successfully"), nil, nil
}

func (h *Handler) ListAssetComments(ctx context.Context, req *mcp.CallToolRequest, args ListAssetCommentsArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetCommentListWithResponse(ctx, h.workspaceID, projectID, assetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list comments: %v", err)), nil, nil
	}
	if resp.JSON200 == nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to list comments: status %d", resp.StatusCode())), nil, nil
	}

	return tool.JSONResult(resp.JSON200), nil, nil
}

func (h *Handler) DeleteAssetComment(ctx context.Context, req *mcp.CallToolRequest, args DeleteAssetCommentArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	commentID, err := id.CommentIDFrom(args.CommentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid comment ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetCommentDeleteWithResponse(ctx, h.workspaceID, projectID, assetID, commentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to delete comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment deleted successfully"), nil, nil
}

func (h *Handler) UpdateAssetComment(ctx context.Context, req *mcp.CallToolRequest, args UpdateAssetCommentArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	commentID, err := id.CommentIDFrom(args.CommentID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid comment ID: %v", err)), nil, nil
	}

	body := cms.AssetCommentUpdateJSONRequestBody{
		Content: &args.Content,
	}

	resp, err := h.client.AssetCommentUpdateWithResponse(ctx, h.workspaceID, projectID, assetID, commentID, body)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to update comment: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to update comment: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Comment updated successfully"), nil, nil
}

func (h *Handler) PublishAsset(ctx context.Context, req *mcp.CallToolRequest, args PublishAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetPublishWithResponse(ctx, h.workspaceID, projectID, assetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to publish asset: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to publish asset: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Asset published successfully"), nil, nil
}

func (h *Handler) UnpublishAsset(ctx context.Context, req *mcp.CallToolRequest, args UnpublishAssetArgs) (*mcp.CallToolResult, any, error) {
	projectID := project.IDOrAlias(args.ProjectID)
	assetID, err := id.AssetIDFrom(args.AssetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Invalid asset ID: %v", err)), nil, nil
	}

	resp, err := h.client.AssetUnpublishWithResponse(ctx, h.workspaceID, projectID, assetID)
	if err != nil {
		return tool.ErrorResult(fmt.Sprintf("Failed to unpublish asset: %v", err)), nil, nil
	}
	if resp.StatusCode() >= 400 {
		return tool.ErrorResult(fmt.Sprintf("Failed to unpublish asset: status %d", resp.StatusCode())), nil, nil
	}

	return tool.TextResult("Asset unpublished successfully"), nil, nil
}
