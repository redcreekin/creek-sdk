package sdk

type ProjectGroupRequest struct {
	ProjectGroupName string `json:"project_group_name" jsonschema:"title=Project Group Name,description=Name of the project group,minLength=1,maxLength=100,required"`
	Description      string `json:"description,omitempty" jsonschema:"title=Description,description=Description of the project group,omitempty"`
	ProjectType      string `json:"project_type" jsonschema:"title=Project Type,description=Type of the project,required"`
}

type ProjectGroupResponse struct {
	ProjectGroupId   string `json:"project_group_id"`
	ProjectGroupName string `json:"project_group_name"`
	Description      string `json:"description"`
	ProjectType      string `json:"project_type"`
	SpaceId          string `json:"space_id"`
}

type ProjectRequest struct {
	ProjectName string                 `json:"project_name"`
	Slug        string                 `json:"slug"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
}

type ProjectResponse struct {
	ProjectId      string                 `json:"project_id"`
	ProjectName    string                 `json:"project_name"`
	Slug           string                 `json:"slug"`
	Description    string                 `json:"description"`
	ProjectType    string                 `json:"project_type"`
	Properties     map[string]interface{} `json:"properties"`
	Workflows      []WorkflowResponse     `json:"workflows"`
	Channels       []ChannelResponse      `json:"channels"`
	Releases       []interface{}          `json:"releases"`
	ProjectGroupId string                 `json:"project_group_id"`
	SpaceId        string                 `json:"space_id"`
	ProjectVersion string                 `json:"project_version"`
	Metadata       map[string]interface{} `json:"metadata"`
}
