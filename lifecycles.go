package sdk

type ReleaseRetentionPolicy struct {
}

type BuildRetentionPolicy struct {
}

type LifecycleStructure struct {
	StructureName              string                 `json:"name" jsonschema:"title=Lifecycle Name,description=Name of the lifecycle,minLength=1,maxLength=100"`
	AutomaticRocketEnvironment []string               `json:"automatic_rockets,omitempty" jsonschema:"title=Automatic Rockets,description=List of automatic rockets for the lifecycle,default=[]"`
	ManualRocketEnvironment    []string               `json:"manual_rockets,omitempty" jsonschema:"title=Manual Rockets,description=List of manual rockets for the lifecycle,default=[]"`
	Stages                     []string               `json:"stages,omitempty" jsonschema:"title=Stages,description=List of stages for the lifecycle,default=[]"`
	ReleaseRetentionPolicy     ReleaseRetentionPolicy `json:"release_retention_policy,omitempty" jsonschema:"title=Release Retention Policy,description=Release Retention policy for the lifecycle,default=null"`
	BuildRetentionPolicy       BuildRetentionPolicy   `json:"build_retention_policy,omitempty" jsonschema:"title=Build Retention Policy,description=Build Retention policy for the lifecycle,default=null"`
}
