package sdk

const (
	WorkflowNameLengthMin = uint64(1)
	WorkflowNameLengthMax = uint64(100)
)

type ActionResponse struct {
	ActionId          string `json:"action_id" jsonschema:"title=Action ID,description=Unique identifier for the action"`
	ActionName        string `json:"action_name" jsonschema:"title=Action Name,description=Name of the action,minLength=1,maxLength=100"`
	ActionDescription string `json:"action_description,omitempty" jsonschema:"title=Action Description,description=Description of the action"`
	ActionType        string `json:"action_type" jsonschema:"title=Action Type,description=Type of the action"`
	ActionTypeVersion string `json:"action_type_version,omitempty" jsonschema:"title=Action Type Version,description=Version of the action type"`
	Status            string `json:"status,omitempty" jsonschema:"title=Status,description=Status of the action,default=active,enum=active,inactive"`
	Position          []int  `json:"position,omitempty" jsonschema:"title=Position,description=Position of the action in the workflow"`
	Properties        JSONB  `json:"properties,omitempty" jsonschema:"title=Properties,description=Properties of the action,default={}"`
}

type WorkflowResponse struct {
	WorkflowId   string           `json:"workflow_id" jsonschema:"title=Workflow ID,description=Unique identifier for the workflow"`
	WorkflowName string           `json:"workflow_name" jsonschema:"title=Workflow Name,description=Name of the workflow"`
	Description  string           `json:"description" jsonschema:"title=Description,description=Description of the workflow"`
	Actions      []ActionResponse `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow"`
}

type ActionRequest struct {
	ActionName        string `json:"action_name" jsonschema:"title=Action Name,description=Name of the action" jsonschema_extras:"order=1"`
	ActionDescription string `json:"action_description,omitempty" jsonschema:"title=Action Description,description=Description of the action" jsonschema_extras:"order=2"`
	ActionType        string `json:"action_type" jsonschema:"title=Action Type,description=Type of the action" jsonschema_extras:"order=3"`
	ActionTypeVersion string `json:"action_type_version" jsonschema:"title=Action Type Version,description=Version of the action type" jsonschema_extras:"order=4"`
	Status            string `json:"status" jsonschema:"title=Status,description=Status of the action" jsonschema_extras:"order=5"`
	Position          []int  `json:"position" jsonschema:"title=Position,description=Position of the action in the workflow"`
	Properties        JSONB  `json:"properties,omitempty" jsonschema:"title=Properties,description=Properties of the action"`
}

type WorkflowRequest struct {
	WorkflowName string          `json:"workflow_name" jsonschema:"title=Workflow Name,description=Name of the workflow" jsonschema_extras:"order=1"`
	Description  string          `json:"description,omitempty" jsonschema:"title=Description,description=Description of the workflow" jsonschema_extras:"order=2"`
	Actions      []ActionRequest `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow" jsonschema_extras:"order=3,mode=edit"`
}
