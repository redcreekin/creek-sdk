package sdk

type EnvironmentRequest struct {
	EnvironmentName       string `json:"environment_name" jsonschema:"title=Environment Name,description=Name of the environment,minLength=1,maxLength=100,required"`
	Description           string `json:"description,omitempty" jsonschema:"title=Description,description=Description of the environment,omitempty"`
	DynamicInfrastructure bool   `json:"dynamic_infrastructure" jsonschema:"title=Dynamic Infrastructure,description=Indicates if the environment uses dynamic infrastructure,default=false"`
	EnableGuidedFailure   bool   `json:"enable_guided_failure" jsonschema:"title=Enable Guided Failure,description=Indicates if guided failure is enabled for the environment,default=false"`
}

type EnvironmentResponse struct {
	EnvironmentId         string `json:"environment_id" jsonschema:"title=Environment Id"`
	EnvironmentName       string `json:"environment_name" jsonschema:"title=Environment Name,description=Name of the environment"`
	Description           string `json:"description,omitempty" jsonschema:"title=Description,description=Description of the environment,default=No description provided"`
	EnvironmentSortOrder  int    `json:"environment_sort_order" jsonschema:"title=Environment Sort Order,description=Sort order of the environment"`
	DynamicInfrastructure bool   `json:"dynamic_infrastructure" jsonschema:"title=Dynamic Infrastructure,description=Indicates if the environment uses dynamic infrastructure,default=false"`
	EnableGuidedFailure   bool   `json:"enable_guided_failure" jsonschema:"title=Enable Guided Failure,description=Indicates if guided failure is enabled for the environment,default=false"`
}
