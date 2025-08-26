package sdk

type PackageList struct {

}
type ReleaseRequest struct {
	ReleaseNumber string `json:"release_number"`
	ReleaseNotes  string `json:"release_notes"`
	LifecycleId   string `json:"lifecycle_id"`
	ChannelId     string `json:"channel_id"`
	ProjectId     string `json:"project_id"`
}

type ReleaseResponse struct {
	ReleaseId     string `json:"release_id"`
	ReleaseNumber string `json:"release_number"`
	ReleaseNotes  string `json:"release_notes"`
	LifecycleId   string `json:"lifecycle_id"`
	ChannelId     string `json:"channel_id"`
	ProjectId     string `json:"project_id"`
	SpaceId       string `json:"space_id"`
}
