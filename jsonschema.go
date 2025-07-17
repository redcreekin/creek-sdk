package sdk

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/invopop/jsonschema"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONB)
		return nil
	}
	return json.Unmarshal(value.([]byte), j)
}

func GetActionJsonSchema(actionModel interface{}) *jsonschema.Schema {
	reflector := jsonschema.Reflector{Anonymous: false}
	actionSchema := reflector.Reflect(actionModel)
	return actionSchema
}

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
	workflowRequestSchema.Title = "Workflow Response"
	workflowRequestSchema.Description = "Schema for workflow response"
	propWorkflowName, _ := workflowRequestSchema.Definitions["WorkflowRequest"].Properties.Get("workflow_name")
	propWorkflowName.Pattern = EntityNamePattern
	propWorkflowName.MinLength = Uint64Ptr(WorkflowNameLengthMin)
	propWorkflowName.MaxLength = Uint64Ptr(WorkflowNameLengthMax)

	propWorkflowActions, _ := workflowRequestSchema.Definitions["WorkflowRequest"].Properties.Get("actions")
	output := propWorkflowActions.(*jsonschema.Schema)

	return workflowRequestSchema
}
