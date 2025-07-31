package sdk

const (
	WorkflowNameLengthMin = uint64(1)
	WorkflowNameLengthMax = uint64(100)
)

type ActionResponse struct {
	ActionId          string `json:"station.action.id" jsonschema:"title=Action ID,description=Unique identifier for the action"`
	ActionName        string `json:"station.action.name" jsonschema:"title=Action Name,description=Name of the action,minLength=1,maxLength=100"`
	ActionDescription string `json:"station.action.description,omitempty" jsonschema:"title=Action Description,description=Description of the action"`
	ActionNotes       string `json:"station.action.notes,omitempty" jsonschema:"title=Action Notes,description=Notes for the action"`
	ActionType        string `json:"station.action.type" jsonschema:"title=Action Type,description=Type of the action"`
	ActionTypeVersion string `json:"station.action.type_version,omitempty" jsonschema:"title=Action Type Version,description=Version of the action type"`

	Status     string          `json:"station.action.status,omitempty" jsonschema:"title=Status,description=Status of the action,default=active,enum=active,inactive"`
	Position   []int           `json:"station.action.position,omitempty" jsonschema:"title=Position,description=Position of the action in the workflow"`
	ConnectTo  []ActionConnect `json:"station.action.connect_to,omitempty" jsonschema:"title=Connect To,description=Actions to connect to,default=[]"`
	Conditions JSONB           `json:"station.action.conditions" jsonschema:"title=Conditions,description=Conditions for the action,default={}" jsonschema_extras:"mode=edit"`
	Properties JSONB           `json:"station.action.properties,omitempty" jsonschema:"title=Properties,description=Properties of the action,default={}" jsonschema_extras:"mode=edit"`
}

type WorkflowResponse struct {
	WorkflowId   string           `json:"station.workflow.id" jsonschema:"title=Workflow ID,description=Unique identifier for the workflow"`
	WorkflowName string           `json:"station.workflow.name" jsonschema:"title=Workflow Name,description=Name of the workflow"`
	Description  string           `json:"station.workflow.description,omitempty" jsonschema:"title=Description,description=Description of the workflow"`
	Actions      []ActionResponse `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow"`
	ProjectId    string           `json:"station.project.id" jsonschema:"title=Project ID,description=Unique identifier for the project"`
	SpaceId      string           `json:"station.space.id" jsonschema:"title=Space ID,description=Unique identifier for the space"`
}

type ActionConnect struct {
	NodeType string `json:"node_type,omitempty" jsonschema:"title=Node Type,description=Type of the node,default=output"`
	NodeId   string `json:"node_id,omitempty" jsonschema:"title=Node ID,description=Unique identifier for the node"`
}

type ActionRequest struct {
	ActionId          string          `json:"station.action.id,omitempty" jsonschema:"title=Action ID,description=Unique identifier for the action" jsonschema_extras:"mode=edit,order=1"`
	ActionName        string          `json:"station.action.name" jsonschema:"title=Action Name,description=Name of the action" jsonschema_extras:"order=2"`
	ActionDescription string          `json:"station.action.description,omitempty" jsonschema:"title=Action Description,description=Description of the action" jsonschema_extras:"order=3"`
	ActionNotes       string          `json:"station.action.notes,omitempty" jsonschema:"title=Action Notes,description=Notes for the action" jsonschema_extras:"mode=edit"`
	ActionType        string          `json:"station.action.type" jsonschema:"title=Action Type,description=Type of the action" jsonschema_extras:"order=4"`
	ActionTypeVersion string          `json:"station.action.type_version" jsonschema:"title=Action Type Version,description=Version of the action type" jsonschema_extras:"order=5"`
	Status            string          `json:"station.action.status" jsonschema:"title=Status,description=Status of the action" jsonschema_extras:"order=6"`
	Position          []int           `json:"station.action.position" jsonschema:"title=Position,description=Position of the action in the workflow"`
	ConnectTo         []ActionConnect `json:"station.action.connect_to,omitempty" jsonschema:"title=Connect To,description=Actions to connect to,default=[]" jsonschema_extras:"mode=edit"`
	Conditions        JSONB           `json:"station.action.conditions" jsonschema:"title=Conditions,description=Conditions for the action,default={}" jsonschema_extras:"mode=edit"`
	Properties        JSONB           `json:"station.action.properties,omitempty" jsonschema:"title=Properties,description=Properties of the action,default={}" jsonschema_extras:"mode=edit"`
}

type WorkflowRequest struct {
	WorkflowName string          `json:"station.workflow.name" jsonschema:"title=Workflow Name,description=Name of the workflow" jsonschema_extras:"order=1"`
	Description  string          `json:"station.workflow.description,omitempty" jsonschema:"title=Description,description=Description of the workflow" jsonschema_extras:"order=2"`
	Actions      []ActionRequest `json:"actions,omitempty" jsonschema:"title=Actions,description=List of actions in the workflow" jsonschema_extras:"order=3,mode=edit"`
	ProjectId    string          `json:"station.workflow.project_id" jsonschema:"title=Project ID,description=Unique identifier for the project" jsonschema_extras:"mode=edit"`
	SpaceId      string          `json:"station.workflow.space_id" jsonschema:"title=Space ID,description=Unique identifier for the space" jsonschema_extras:"mode=edit"`
}
