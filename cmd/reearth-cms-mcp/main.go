package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/reearth/reearth-cms-mcp/internal/handler"
	"github.com/reearth/reearth-cms-mcp/internal/logging"
	"github.com/reearth/reearth-cms-mcp/internal/version"
)

const defaultBaseURL = "https://api.cms.reearth.io/api"

func main() {
	cfg := handler.Config{
		Token:       os.Getenv("REEARTH_CMS_TOKEN"),
		BaseURL:     os.Getenv("REEARTH_CMS_BASE_URL"),
		WorkspaceID: os.Getenv("REEARTH_CMS_WORKSPACE_ID"),
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}

	h, err := handler.New(cfg)
	if err != nil {
		log.Fatalf("failed to create handler: %v", err)
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    version.Name,
		Title:   version.Title,
		Version: version.Version,
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "comment_to_asset",
		Description: "Add a comment to an asset in Re:Earth CMS",
	}, h.CommentToAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "comment_to_item",
		Description: "Add a comment to an item in Re:Earth CMS",
	}, h.CommentToItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "copy_model",
		Description: "Copy an existing model in Re:Earth CMS",
	}, h.CopyModel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_asset",
		Description: "Create a new asset from a URL in Re:Earth CMS",
	}, h.CreateAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_field",
		Description: "Create a new field in a schema",
	}, h.CreateField)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_group",
		Description: "Create a new group in Re:Earth CMS",
	}, h.CreateGroup)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_item",
		Description: "Create a new item in Re:Earth CMS",
	}, h.CreateItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_model",
		Description: "Create a new model in Re:Earth CMS",
	}, h.CreateModel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_asset",
		Description: "Delete an asset from Re:Earth CMS",
	}, h.DeleteAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_asset_comment",
		Description: "Delete a comment from an asset in Re:Earth CMS",
	}, h.DeleteAssetComment)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_field",
		Description: "Delete a field from a schema",
	}, h.DeleteField)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_group",
		Description: "Delete a group from Re:Earth CMS",
	}, h.DeleteGroup)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_item",
		Description: "Delete an item from Re:Earth CMS",
	}, h.DeleteItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_item_comment",
		Description: "Delete a comment from an item in Re:Earth CMS",
	}, h.DeleteItemComment)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_model",
		Description: "Delete a model from Re:Earth CMS",
	}, h.DeleteModel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_asset",
		Description: "Get a specific asset by ID from Re:Earth CMS",
	}, h.GetAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_group",
		Description: "Get a specific group by ID from Re:Earth CMS",
	}, h.GetGroup)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_item",
		Description: "Get a specific item by ID from Re:Earth CMS",
	}, h.GetItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_items_as_csv",
		Description: "Export items as CSV from Re:Earth CMS",
	}, h.GetItemsAsCSV)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_items_as_geojson",
		Description: "Export items as GeoJSON from Re:Earth CMS",
	}, h.GetItemsAsGeoJSON)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_model",
		Description: "Get a specific model by ID from Re:Earth CMS",
	}, h.GetModel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_model_metadata_schema",
		Description: "Get the metadata schema for a model",
	}, h.GetModelMetadataSchema)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_model_schema",
		Description: "Get the JSON schema for a model",
	}, h.GetModelSchema)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_project",
		Description: "Get a project by ID from Re:Earth CMS",
	}, h.GetProject)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_asset_comments",
		Description: "List comments on an asset in Re:Earth CMS",
	}, h.ListAssetComments)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_assets",
		Description: "List assets in a Re:Earth CMS project",
	}, h.ListAssets)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_groups",
		Description: "List all groups in a Re:Earth CMS project",
	}, h.ListGroups)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_item_comments",
		Description: "List comments on an item in Re:Earth CMS",
	}, h.ListItemComments)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_items",
		Description: "List items in a Re:Earth CMS project model",
	}, h.ListItems)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_models",
		Description: "List all models in a Re:Earth CMS project",
	}, h.ListModels)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_projects",
		Description: "List all projects in a workspace",
	}, h.ListProjects)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "publish_asset",
		Description: "Publish an asset in Re:Earth CMS",
	}, h.PublishAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "publish_item",
		Description: "Publish an item in Re:Earth CMS",
	}, h.PublishItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "unpublish_asset",
		Description: "Unpublish an asset in Re:Earth CMS",
	}, h.UnpublishAsset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_asset_comment",
		Description: "Update a comment on an asset in Re:Earth CMS",
	}, h.UpdateAssetComment)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_field",
		Description: "Update an existing field in a schema",
	}, h.UpdateField)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_group",
		Description: "Update an existing group in Re:Earth CMS",
	}, h.UpdateGroup)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_item",
		Description: "Update an existing item in Re:Earth CMS",
	}, h.UpdateItem)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_item_comment",
		Description: "Update a comment on an item in Re:Earth CMS",
	}, h.UpdateItemComment)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_model",
		Description: "Update an existing model in Re:Earth CMS",
	}, h.UpdateModel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_project",
		Description: "Update a project in Re:Earth CMS",
	}, h.UpdateProject)

	var transport mcp.Transport = &mcp.StdioTransport{}
	if logging.IsDebugEnabled() {
		transport = logging.NewLoggingTransport(transport, os.Stderr)
	}

	if err := server.Run(context.Background(), transport); err != nil {
		log.Fatal(err)
	}
}
