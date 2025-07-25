package sdk

const (
	WorkflowNameLengthMin = uint64(1)
	WorkflowNameLengthMax = uint64(100)
)

type ActionResponse struct {
	ActionId          string `json:"station.action.id" jsonschema:"title=Action ID,description=Unique identifier for the action"`
	ActionName        string `json:"station.action.name" jsonschema:"title=Action Name,description=Name of the action,minLength=1,maxLength=100"`
	ActionDescription string `json:"station.action.description,omitempty" jsonschema:"title=Action Description,description=Description of the action"`
	ActionType        string `json:"station.action.type" jsonschema:"title=Action Type,description=Type of the action"`
	ActionTypeVersion string `json:"station.action.type_version,omitempty" jsonschema:"title=Action Type Version,description=Version of the action type"`
	Status            string `json:"station.action.status,omitempty" jsonschema:"title=Status,description=Status of the action,default=active,enum=active,inactive"`
	Position          []int  `json:"station.action.position,omitempty" jsonschema:"title=Position,description=Position of the action in the workflow"`
	Conditions        JSONB  `json:"station.action.conditions,omitempty" jsonschema:"title=Conditions,description=Conditions for the action,default={}" jsonschema_extras:"mode=edit"`
	Properties        JSONB  `json:"station.action.properties,omitempty" jsonschema:"title=Properties,description=Properties of the action,default={}" jsonschema_extras:"mode=edit"`
}

type WorkflowResponse struct {
	WorkflowId   string           `json:"workflow_id" jsonschema:"title=Workflow ID,description=Unique identifier for the workflow"`
	WorkflowName string           `json:"workflow_name" jsonschema:"title=Workflow Name,description=Name of the workflow"`
	Description  string           `json:"description" jsonschema:"title=Description,description=Description of the workflow"`
	Actions      []ActionResponse `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow"`
}

type ActionCondition struct {
	Environments struct {
		All     bool     `json:"all" jsonschema:"title=All,description=Whether all environments must match,default=true"`
		Include []string `json:"include,omitempty" jsonschema:"title=Include,description=List of environments to include"`
		Exclude []string `json:"exclude,omitempty" jsonschema:"title=Exclude,description=List of environments to exclude"`
	} `json:"station.action.runon.environments" jsonschema:"title=Environments,description=Conditions for environments"`
	Option   string `json:"station.action.runon.option,omitempty" jsonschema:"title=Option,description=Run on option,default=all" jsonschema_extras:"default=always"`
	Trigger  string `json:"station.action.runon.trigger,omitempty" jsonschema:"title=Trigger,description=Trigger for the action,default=manual" jsonschema_extras:"default=wait_previous_step"`
	Required bool   `json:"station.action.runon.required,omitempty" jsonschema:"title=Required,description=Whether the action is required,default=false" jsonschema_extras:"default=false"`
}

type ActionRequest struct {
	ActionId          string          `json:"station.action.id,omitempty" jsonschema:"title=Action ID,description=Unique identifier for the action" jsonschema_extras:"mode=edit,order=1"`
	ActionName        string          `json:"station.action.name" jsonschema:"title=Action Name,description=Name of the action" jsonschema_extras:"order=2"`
	ActionDescription string          `json:"station.action.description,omitempty" jsonschema:"title=Action Description,description=Description of the action" jsonschema_extras:"order=3"`
	ActionType        string          `json:"station.action.type" jsonschema:"title=Action Type,description=Type of the action" jsonschema_extras:"order=4"`
	ActionTypeVersion string          `json:"station.action.type_version" jsonschema:"title=Action Type Version,description=Version of the action type" jsonschema_extras:"order=5"`
	Status            string          `json:"station.action.status" jsonschema:"title=Status,description=Status of the action" jsonschema_extras:"order=6"`
	Position          []int           `json:"station.action.position" jsonschema:"title=Position,description=Position of the action in the workflow"`
	ConnectTo         []string        `json:"station.action.connect_to,omitempty" jsonschema:"title=Connect To,description=Actions to connect to,default=[]" jsonschema_extras:"mode=edit"`
	Conditions        ActionCondition `json:"station.action.conditions,omitempty" jsonschema:"title=Conditions,description=Conditions for the action,default={}" jsonschema_extras:"mode=edit"`
	Properties        JSONB           `json:"station.action.properties,omitempty" jsonschema:"title=Properties,description=Properties of the action,default={}" jsonschema_extras:"mode=edit"`
}

type WorkflowRequest struct {
	WorkflowName string          `json:"station.workflow.name" jsonschema:"title=Workflow Name,description=Name of the workflow" jsonschema_extras:"order=1"`
	Description  string          `json:"station.workflow.description,omitempty" jsonschema:"title=Description,description=Description of the workflow" jsonschema_extras:"order=2"`
	Actions      []ActionRequest `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow" jsonschema_extras:"order=3,mode=edit"`
}
