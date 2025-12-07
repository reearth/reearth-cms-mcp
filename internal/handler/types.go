package handler

type FieldArg struct {
	Key   string `json:"key" jsonschema:"description=The field key"`
	Type  string `json:"type" jsonschema:"description=The field type (text, number, bool, date, asset, reference, etc.)"`
	Value any    `json:"value" jsonschema:"description=The field value"`
}

type GetAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
}

type ListAssetsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"description=Search keyword"`
}

type DeleteAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
}

type CreateAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	URL       string `json:"url" jsonschema:"description=The URL of the asset to create,required"`
}

type CommentToAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
	Content   string `json:"content" jsonschema:"description=The comment content,required"`
}

type ListAssetCommentsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
}

type DeleteAssetCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
	CommentID string `json:"comment_id" jsonschema:"description=The comment ID,required"`
}

type UpdateAssetCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
	CommentID string `json:"comment_id" jsonschema:"description=The comment ID,required"`
	Content   string `json:"content" jsonschema:"description=The updated comment content,required"`
}

type PublishAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
}

type UnpublishAssetArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	AssetID   string `json:"asset_id" jsonschema:"description=The asset ID,required"`
}

type GetItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
}

type ListItemsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"description=Search keyword"`
}

type CreateItemArgs struct {
	ProjectID string     `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string     `json:"model_id" jsonschema:"description=The model ID or key,required"`
	Fields    []FieldArg `json:"fields" jsonschema:"description=The fields for the item,required"`
}

type UpdateItemArgs struct {
	ProjectID string     `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string     `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string     `json:"item_id" jsonschema:"description=The item ID,required"`
	Fields    []FieldArg `json:"fields" jsonschema:"description=The fields to update,required"`
}

type DeleteItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
}

type CommentToItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
	Content   string `json:"content" jsonschema:"description=The comment content,required"`
}

type ListItemCommentsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
}

type DeleteItemCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
	CommentID string `json:"comment_id" jsonschema:"description=The comment ID,required"`
}

type UpdateItemCommentArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
	CommentID string `json:"comment_id" jsonschema:"description=The comment ID,required"`
	Content   string `json:"content" jsonschema:"description=The updated comment content,required"`
}

type PublishItemArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	ItemID    string `json:"item_id" jsonschema:"description=The item ID,required"`
}

type GetItemsAsCSVArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetItemsAsGeoJSONArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type GetModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
}

type ListModelsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
	Keyword   string `json:"keyword,omitempty" jsonschema:"description=Search keyword"`
}

type CreateModelArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"description=The project ID,required"`
	Name        string `json:"name" jsonschema:"description=The model name,required"`
	Key         string `json:"key,omitempty" jsonschema:"description=The model key"`
	Description string `json:"description,omitempty" jsonschema:"description=The model description"`
}

type UpdateModelArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID     string `json:"model_id" jsonschema:"description=The model ID or key,required"`
	Name        string `json:"name,omitempty" jsonschema:"description=The model name"`
	Key         string `json:"key,omitempty" jsonschema:"description=The model key"`
	Description string `json:"description,omitempty" jsonschema:"description=The model description"`
}

type DeleteModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
}

type CopyModelArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The source model ID or key,required"`
	Name      string `json:"name,omitempty" jsonschema:"description=The name for the copied model"`
	Key       string `json:"key,omitempty" jsonschema:"description=The key for the copied model"`
}

type GetModelSchemaArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
}

type GetModelMetadataSchemaArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	ModelID   string `json:"model_id" jsonschema:"description=The model ID or key,required"`
}

type GetProjectArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
}

type ListProjectsArgs struct {
	Page    int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
	Keyword string `json:"keyword,omitempty" jsonschema:"description=Search keyword"`
}

type UpdateProjectArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"description=The project ID,required"`
	Name        string `json:"name,omitempty" jsonschema:"description=The project name"`
	Description string `json:"description,omitempty" jsonschema:"description=The project description"`
	Alias       string `json:"alias,omitempty" jsonschema:"description=The project alias"`
}

type GetGroupArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	GroupID   string `json:"group_id" jsonschema:"description=The group ID or key,required"`
}

type ListGroupsArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	Page      int    `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"description=Number of items per page"`
}

type CreateGroupArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"description=The project ID,required"`
	Name        string `json:"name" jsonschema:"description=The group name,required"`
	Key         string `json:"key" jsonschema:"description=The group key,required"`
	Description string `json:"description,omitempty" jsonschema:"description=The group description"`
}

type UpdateGroupArgs struct {
	ProjectID   string `json:"project_id" jsonschema:"description=The project ID,required"`
	GroupID     string `json:"group_id" jsonschema:"description=The group ID or key,required"`
	Name        string `json:"name,omitempty" jsonschema:"description=The group name"`
	Key         string `json:"key,omitempty" jsonschema:"description=The group key"`
	Description string `json:"description,omitempty" jsonschema:"description=The group description"`
}

type DeleteGroupArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	GroupID   string `json:"group_id" jsonschema:"description=The group ID or key,required"`
}

type CreateFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	SchemaID  string `json:"schema_id" jsonschema:"description=The schema ID,required"`
	Key       string `json:"key" jsonschema:"description=The field key,required"`
	Type      string `json:"type" jsonschema:"description=The field type (text, number, bool, date, asset, reference, etc.),required"`
	Multiple  bool   `json:"multiple,omitempty" jsonschema:"description=Whether the field accepts multiple values"`
	Required  bool   `json:"required,omitempty" jsonschema:"description=Whether the field is required"`
}

type UpdateFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	SchemaID  string `json:"schema_id" jsonschema:"description=The schema ID,required"`
	FieldID   string `json:"field_id" jsonschema:"description=The field ID or key,required"`
	Key       string `json:"key,omitempty" jsonschema:"description=The field key"`
	Type      string `json:"type,omitempty" jsonschema:"description=The field type"`
	Multiple  *bool  `json:"multiple,omitempty" jsonschema:"description=Whether the field accepts multiple values"`
	Required  *bool  `json:"required,omitempty" jsonschema:"description=Whether the field is required"`
}

type DeleteFieldArgs struct {
	ProjectID string `json:"project_id" jsonschema:"description=The project ID,required"`
	SchemaID  string `json:"schema_id" jsonschema:"description=The schema ID,required"`
	FieldID   string `json:"field_id" jsonschema:"description=The field ID or key,required"`
}
