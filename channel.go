package sdk

type ChannelRequest struct {
	ChannelName string         `json:"channel_name"`
	Description string         `json:"description"`
	IsDefault   bool           `json:"is_default"`
	Rules       []ChannelRules `json:"rules"`
	LifecycleId string         `json:"lifecycle_id"`
}

type ChannelRules struct {
	RuleId         string                 `json:"rule_id"`
	VersionRange   string                 `json:"version_range"`
	Tag            string                 `json:"tag"`
	ActionPackages []map[string]string    `json:"action_packages"`
	Links          map[string]interface{} `json:"links"`
	Actions        []string               `json:"actions"`
}
type ChannelResponse struct {
	ChannelId   string         `json:"channel_id"`
	ChannelName string         `json:"channel_name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	IsDefault   bool           `json:"is_default"`
	Rules       []ChannelRules `json:"rules"`
	LifecycleId string         `json:"lifecycle_id"`
}
