package sdk

type ProjectGroupRequest struct {
	ProjectGroupName string `json:"project_group_name" jsonschema:"title=Project Group Name,description=Name of the project group,minLength=1,maxLength=100,required"`
	Description      string `json:"description,omitempty" jsonschema:"title=Description,description=Description of the project group,omitempty"`
	ProjectType      string `json:"project_type" jsonschema:"title=Project Type,description=Type of the project,required"`
}
