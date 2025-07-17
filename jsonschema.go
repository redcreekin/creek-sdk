package sdk

import (
	"github.com/invopop/jsonschema"
)

func GetEnvironmentRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	environmentSchema := reflector.Reflect(EnvironmentRequest{})
	environmentSchema.Title = "Environment Request"
	environmentSchema.Description = "Schema for environment request"
	return environmentSchema
}

func GetWorkflowRequestJsonSchema() *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	workflowRequestSchema := reflector.Reflect(&WorkflowRequest{})
	actionRequestSchema := reflector.Reflect(&ActionRequest{})

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
