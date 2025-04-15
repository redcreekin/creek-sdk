package slog

import "github.com/rockbears/log"

const (
	Action         = log.Field("action")
	Administrator  = log.Field("administrator")
	AuthConsumerId = log.Field("auth_consumer_id")
	AuthSessionIAT = log.Field("auth_session_iat")
	AuthSessionID  = log.Field("auth_session_id")
	AuthUserID     = log.Field("auth_user_id")
	AuthUsername   = log.Field("auth_user_name")
	Deprecated     = log.Field("deprecated")
	Duration       = log.Field("duration_milliseconds_num")
	Goroutine      = log.Field("goroutine")
	IPAddress      = log.Field("ip_address")
	Latency        = log.Field("latency")
	LatencyNum     = log.Field("latency_num")
	Method         = log.Field("method")
	GpgKey         = log.Field("gpg_key")
	RbacRole       = log.Field("rbac_role")
	RequestID      = log.Field("request_id")
	RequestURI     = log.Field("request_uri")

	Operation     = log.Field("operation")
	Route         = log.Field("route")
	Size          = log.Field("size_num")
	Stacktrace    = log.Field("stack_trace")
	Sudo          = log.Field("sudo")
	Workflow      = log.Field("workflow")
	WorkflowRunID = log.Field("workflow_run_id")
	Component     = log.Field("component")
	Project       = log.Field("project")
)

func init() {
	log.RegisterField(
		Action,
		Administrator,
		AuthConsumerId,
		AuthSessionIAT,
		AuthSessionID,
		AuthUserID,
		AuthUsername,
		Deprecated,
		Duration,
		Goroutine,
		IPAddress,
		Latency,
		LatencyNum,
		Method,
		GpgKey,
		RbacRole,
		RequestID,
		RequestURI,
		Operation,
		Route,
		Size,
		Stacktrace,
		Sudo,
		Workflow,
		WorkflowRunID,
		Component,
		Project,
	)
}
