package sdk

import (
	"github.com/invopop/jsonschema"
)

func GetEnvironmentRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	environmentSchema := reflector.Reflect(EnvironmentRequest{})
	environmentSchema.Title = "Environment Request"
	environmentSchema.Description = "Schema for environment request"
	environmentName, _ := environmentSchema.Definitions["EnvironmentRequest"].Properties.Get("environment_name")
	environmentName.Pattern = EntityNamePattern
	return environmentSchema
}

func GetWorkflowRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	workflowRequestSchema := reflector.Reflect(&WorkflowRequest{})
	actionRequestSchema := GetActionRequestJsonSchema()

	workflowRequestSchema.Title = "Workflow Request"
	workflowRequestSchema.Description = "Schema for workflow response"
	propWorkflowName, _ := workflowRequestSchema.Definitions["WorkflowRequest"].Properties.Get("workflow_name")
	propWorkflowName.Pattern = EntityNamePattern
	propWorkflowName.MinLength = Uint64Ptr(WorkflowNameLengthMin)
	propWorkflowName.MaxLength = Uint64Ptr(WorkflowNameLengthMax)

	propActions, _ := workflowRequestSchema.Definitions["WorkflowRequest"].Properties.Get("actions")
	if propActions.Definitions == nil {
		propActions.Definitions = actionRequestSchema.Definitions
	}

	return workflowRequestSchema
}

func GetActionRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	actionRequestSchema := reflector.Reflect(&ActionRequest{})
	actionRequestSchema.Title = "Action Request"
	actionRequestSchema.Description = "Schema for action request"
	propActionName, _ := actionRequestSchema.Definitions["ActionRequest"].Properties.Get("station.action.name")
	propActionName.Pattern = EntityNamePattern

	propActionId, _ := actionRequestSchema.Definitions["ActionRequest"].Properties.Get("station.action.id")
	propActionId.Pattern = EntityIdPattern

	return actionRequestSchema
}

func GetProjectGroupRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	groupRequestSchema := reflector.Reflect(&ProjectGroupRequest{})
	groupRequestSchema.Title = "Project Group Request"
	groupRequestSchema.Description = "Schema for project group request"
	projectGroupName, _ := groupRequestSchema.Definitions["ProjectGroupRequest"].Properties.Get("project_group_name")
	projectGroupName.Pattern = EntityNamePattern
	projectType, _ := groupRequestSchema.Definitions["ProjectGroupRequest"].Properties.Get("project_type")
	projectType.Enum = []any{ProjectTypeBuild, ProjectTypeDeploy, ProjectTypeData, ProjectTypeAI, ProjectTypeTest, ProjectTypeSecurity}
	return groupRequestSchema
}
