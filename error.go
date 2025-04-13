package creek_sdk

import (
	"fmt"
	"net/http"
)

var (
	ErrUnknownError              = Error{Code: 1, Status: http.StatusInternalServerError}
	ErrProcessAlreadyUpdate      = Error{Code: 2, Status: http.StatusBadRequest}
	ErrActionAlreadyUpdate       = Error{Code: 3, Status: http.StatusBadRequest}
	ErrActionNotFound            = Error{Code: 4, Status: http.StatusNotFound}
	ErrActionLoop                = Error{Code: 5, Status: http.StatusBadRequest}
	ErrInvalidCode               = Error{Code: 6, Status: http.StatusBadRequest}
	ErrInvalidProject            = Error{Code: 7, Status: http.StatusBadRequest}
	ErrInvalidProjectGroup       = Error{Code: 8, Status: http.StatusBadRequest}
	ErrProjectHasWorkflow        = Error{Code: 9, Status: http.StatusForbidden}
	ErrProjectGroupHasProject    = Error{Code: 10, Status: http.StatusForbidden}
	ErrUnauthorized              = Error{Code: 11, Status: http.StatusUnauthorized}
	ErrForbidden                 = Error{Code: 12, Status: http.StatusForbidden}
	ErrWorkflowNotFound          = Error{Code: 13, Status: http.StatusBadRequest}
	ErrWorkflowNotAttached       = Error{Code: 14, Status: http.StatusBadRequest}
	ErrNoEnvironmentProvided     = Error{Code: 15, Status: http.StatusBadRequest}
	ErrEnvironmentProvided       = Error{Code: 16, Status: http.StatusBadRequest}
	ErrUnknownEnv                = Error{Code: 17, Status: http.StatusBadRequest}
	ErrEnvironmentExist          = Error{Code: 18, Status: http.StatusForbidden}
	ErrNoWorkflowProcess         = Error{Code: 19, Status: http.StatusNotFound}
	ErrInvalidUsername           = Error{Code: 21, Status: http.StatusBadRequest}
	ErrInvalidEmail              = Error{Code: 22, Status: http.StatusBadRequest}
	ErrGroupPresent              = Error{Code: 23, Status: http.StatusBadRequest}
	ErrInvalidName               = Error{Code: 24, Status: http.StatusBadRequest}
	ErrInvalidUser               = Error{Code: 25, Status: http.StatusBadRequest}
	ErrBuildArchived             = Error{Code: 26, Status: http.StatusBadRequest}
	ErrNoEnvironment             = Error{Code: 27, Status: http.StatusNotFound}
	ErrModelNameExist            = Error{Code: 28, Status: http.StatusForbidden}
	ErrNoProject                 = Error{Code: 30, Status: http.StatusNotFound}
	ErrVariableExists            = Error{Code: 31, Status: http.StatusForbidden}
	ErrInvalidGroupPattern       = Error{Code: 32, Status: http.StatusBadRequest}
	ErrGroupExists               = Error{Code: 33, Status: http.StatusForbidden}
	ErrNotEnoughAdmin            = Error{Code: 34, Status: http.StatusBadRequest}
	ErrInvalidProjectName        = Error{Code: 35, Status: http.StatusBadRequest}
	ErrInvalidApplicationPattern = Error{Code: 36, Status: http.StatusBadRequest}
	ErrInvalidPipelinePattern    = Error{Code: 37, Status: http.StatusBadRequest}
	ErrNotFound                  = Error{Code: 38, Status: http.StatusNotFound}
	ErrInvalidGoPath             = Error{Code: 48, Status: http.StatusBadRequest}
	ErrAlreadyTaken              = Error{Code: 91, Status: http.StatusGone}
	ErrRocketNodeNotFound        = Error{Code: 93, Status: http.StatusNotFound}
	ErrRocketInvalid             = Error{Code: 96, Status: http.StatusBadRequest}
	ErrNotImplemented            = Error{Code: 99, Status: http.StatusNotImplemented}
	ErrInvalidKeyPattern         = Error{Code: 102, Status: http.StatusBadRequest}
	ErrTokenNotFound             = Error{Code: 121, Status: http.StatusNotFound}
	ErrUnsupportedOSArchPlugin   = Error{Code: 129, Status: http.StatusNotFound}
	ErrWorkflowImport            = Error{Code: 140, Status: http.StatusBadRequest}
	ErrInvalidData               = Error{Code: 149, Status: http.StatusBadRequest}
	ErrInvalidGroupAdmin         = Error{Code: 150, Status: http.StatusForbidden}
	ErrInvalidGroupMember        = Error{Code: 151, Status: http.StatusForbidden}
	ErrLocked                    = Error{Code: 164, Status: http.StatusConflict}
	ErrInvalidPassword           = Error{Code: 175, Status: http.StatusBadRequest}
	ErrSignupDisabled            = Error{Code: 182, Status: http.StatusForbidden}
	ErrMFARequired               = Error{Code: 194, Status: http.StatusForbidden}
)

var errorsEnglish = map[int]string{
	ErrUnknownError.Code:           "internal server error",
	ErrProcessAlreadyUpdate.Code:   "process status already updated",
	ErrActionAlreadyUpdate.Code:    "action status already updated",
	ErrActionNotFound.Code:         "action not found",
	ErrActionLoop.Code:             "action definition contains a recursive loop",
	ErrInvalidCode.Code:            "code must be an integer",
	ErrInvalidProject.Code:         "invalid project",
	ErrInvalidProjectGroup.Code:    "invalid project group",
	ErrProjectHasWorkflow.Code:     "project has workflow",
	ErrProjectGroupHasProject.Code: "project group has project",
	ErrUnauthorized.Code:           "unauthorized",
	ErrForbidden.Code:              "forbidden",
	ErrWorkflowNotFound.Code:       "workflow not found",
	ErrWorkflowNotAttached.Code:    "workflow not attached",
}

type Error struct {
	Code       int         `json:"code"`
	Status     int         `json:"-"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	RequestID  string      `json:"request_id,omitempty"`
	StackTrace string      `json:"stack_trace,omitempty"`
	From       string      `json:"from,omitempty"`
}

func (e Error) Error() string {
	var message string
	if e.Message != "" {
		message = e.Message
	} else if en, ok := errorsEnglish[e.Code]; ok {
		message = en
	} else {
		message = errorsEnglish[ErrUnknownError.Code]
	}
	if e.From != "" {
		if e.RequestID != "" {
			message = fmt.Sprintf("%s (from: %s, request_id: %s)", message, e.From, e.RequestID)
		} else {
			message = fmt.Sprintf("%s (from: %s)", message, e.From)
		}
		message = e.From + ": " + message
	} else {
		if e.RequestID != "" {
			message = fmt.Sprintf("%s (request_id: %s)", message, e.RequestID)
		}
		message = "error: " + message
	}
	return message
}

type errorWithStack struct {
	root      error // root error should be wrapped with stack
	httpError Error
}

func (w errorWithStack) Error() string {
	var cause string
	root := w.root.Error()
	if root != "" && root != w.httpError.From && root != w.httpError.Error() {
		cause = fmt.Sprintf("(caused by: %s)", w.root)
	}
	if cause == "" {
		return w.httpError.Error()
	}
	return fmt.Sprintf("%s %s", w.httpError, cause)
}
