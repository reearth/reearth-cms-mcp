package handler

type FieldArg struct {
	Key   string `json:"key" jsonschema:"The field key"`
	Type  string `json:"type" jsonschema:"The field type (text, number, bool, date, asset, reference, etc.)"`
	Value any    `json:"value" jsonschema:"The field value"`
}

type GetAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
}

type ListAssetsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"Search keyword"`
}

type DeleteAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
}

type CreateAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	URL       string `json:"url" jsonschema:"The URL of the asset to create"`
}

type CommentToAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
	Content   string `json:"content" jsonschema:"The comment content"`
}

type ListAssetCommentsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
}

type DeleteAssetCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
	CommentID string `json:"comment_id" jsonschema:"The comment ID"`
}

type UpdateAssetCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
	CommentID string `json:"comment_id" jsonschema:"The comment ID"`
	Content   string `json:"content" jsonschema:"The updated comment content"`
}

type PublishAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
}

type UnpublishAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	AssetID   string `json:"asset_id" jsonschema:"The asset ID"`
}

type GetItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
}

type ListItemsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"Search keyword"`
}

type CreateItemArgs struct {
	ProjectID string     `json:"project_id" jsonschema:"The project ID"`
	ModelID   string     `json:"model_id" jsonschema:"The model ID or key"`
	Fields    []FieldArg `json:"fields" jsonschema:"The fields for the item"`
}

type UpdateItemArgs struct {
	ProjectID string     `json:"project_id" jsonschema:"The project ID"`
	ModelID   string     `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string     `json:"item_id" jsonschema:"The item ID"`
	Fields    []FieldArg `json:"fields" jsonschema:"The fields to update"`
}

type DeleteItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
}

type CommentToItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
	Content   string `json:"content" jsonschema:"The comment content"`
}

type ListItemCommentsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
}

type DeleteItemCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
	CommentID string `json:"comment_id" jsonschema:"The comment ID"`
}

type UpdateItemCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
	CommentID string `json:"comment_id" jsonschema:"The comment ID"`
	Content   string `json:"content" jsonschema:"The updated comment content"`
}

type PublishItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	ItemID    string `json:"item_id" jsonschema:"The item ID"`
}

type GetItemsAsCSVArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
}

type GetItemsAsGeoJSONArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
}

type GetModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
}

type ListModelsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"Search keyword"`
}

type CreateModelArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"The project ID"`
	Name        string `json:"name" jsonschema:"The model name"`
	Key         string `json:"key,omitempty" jsonschema:"The model key"`
	Description string `json:"description,omitempty" jsonschema:"The model description"`
}

type UpdateModelArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"The project ID"`
	ModelID     string `json:"model_id" jsonschema:"The model ID or key"`
	Name        string `json:"name,omitempty" jsonschema:"The model name"`
	Key         string `json:"key,omitempty" jsonschema:"The model key"`
	Description string `json:"description,omitempty" jsonschema:"The model description"`
}

type DeleteModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
}

type CopyModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The source model ID or key"`
	Name      string `json:"name,omitempty" jsonschema:"The name for the copied model"`
	Key       string `json:"key,omitempty" jsonschema:"The key for the copied model"`
}

type GetModelSchemaArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
}

type GetModelMetadataSchemaArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	ModelID   string `json:"model_id" jsonschema:"The model ID or key"`
}

type GetProjectArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
}

type ListProjectsArgs struct {
	Page    int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
	Keyword string `json:"keyword,omitempty" jsonschema:"Search keyword"`
}

type UpdateProjectArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"The project ID"`
	Name        string `json:"name,omitempty" jsonschema:"The project name"`
	Description string `json:"description,omitempty" jsonschema:"The project description"`
	Alias       string `json:"alias,omitempty" jsonschema:"The project alias"`
}

type GetGroupArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	GroupID   string `json:"group_id" jsonschema:"The group ID or key"`
}

type ListGroupsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	Page      int    `json:"page,omitempty" jsonschema:"Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of items per page"`
}

type CreateGroupArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"The project ID"`
	Name        string `json:"name" jsonschema:"The group name"`
	Key         string `json:"key" jsonschema:"The group key"`
	Description string `json:"description,omitempty" jsonschema:"The group description"`
}

type UpdateGroupArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"The project ID"`
	GroupID     string `json:"group_id" jsonschema:"The group ID or key"`
	Name        string `json:"name,omitempty" jsonschema:"The group name"`
	Key         string `json:"key,omitempty" jsonschema:"The group key"`
	Description string `json:"description,omitempty" jsonschema:"The group description"`
}

type DeleteGroupArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	GroupID   string `json:"group_id" jsonschema:"The group ID or key"`
}

type CreateFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	SchemaID  string `json:"schema_id" jsonschema:"The schema ID"`
	Key       string `json:"key" jsonschema:"The field key"`
	Type      string `json:"type" jsonschema:"The field type (text, number, bool, date, asset, reference, etc.)"`
	Multiple  bool   `json:"multiple,omitempty" jsonschema:"Whether the field accepts multiple values"`
	Required  bool   `json:"required,omitempty" jsonschema:"Whether the field is required"`
}

type UpdateFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	SchemaID  string `json:"schema_id" jsonschema:"The schema ID"`
	FieldID   string `json:"field_id" jsonschema:"The field ID or key"`
	Key       string `json:"key,omitempty" jsonschema:"The field key"`
	Type      string `json:"type,omitempty" jsonschema:"The field type"`
	Multiple  *bool  `json:"multiple,omitempty" jsonschema:"Whether the field accepts multiple values"`
	Required  *bool  `json:"required,omitempty" jsonschema:"Whether the field is required"`
}

type DeleteFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"The project ID"`
	SchemaID  string `json:"schema_id" jsonschema:"The schema ID"`
	FieldID   string `json:"field_id" jsonschema:"The field ID or key"`
}
