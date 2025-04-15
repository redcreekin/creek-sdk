package sdk

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rockbears/log"
	"net/http"
	"os"
	"strings"
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
type MultiError []error

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

func (e Error) Translate() string {
	msg, ok := errorsEnglish[e.Code]
	if !ok {
		return errorsEnglish[ErrUnknownError.Code]
	}
	return msg
}

func (e Error) printLight() string {
	var message string
	if e.Message != "" {
		message = e.Message
	} else if en, ok := errorsEnglish[e.Code]; ok {
		message = en
	} else {
		message = errorsEnglish[ErrUnknownError.Code]
	}
	if e.From != "" {
		message = fmt.Sprintf("%s: %s", message, e.From)
	}
	return message
}

func NewError(httpError Error, err error) error {
	// if the given error is nil do nothing
	if err == nil {
		return nil
	}

	// if it's already an error with stack, override the http error and set from value with err cause
	if e, ok := err.(errorWithStack); ok {
		httpError.From = e.httpError.From
		e.httpError = httpError
		return e
	}

	if e, ok := err.(*MultiError); ok {
		var ss []string
		for i := range *e {
			ss = append(ss, ExtractHTTPError((*e)[i]).printLight())
		}
		httpError.From = strings.Join(ss, ", ")
	} else {
		httpError.From = err.Error()
	}

	// if it's a library error create a new error with stack
	return errorWithStack{
		root:      errors.WithStack(err),
		httpError: httpError,
	}
}

func (w errorWithStack) StackTrace() errors.StackTrace {
	errStackTrace, ok := w.root.(log.StackTracer)
	if ok {
		return errStackTrace.StackTrace()
	}
	return nil
}

func ExtractHTTPError(source error) Error {
	var httpError Error

	// try to recognize http error from source
	switch e := source.(type) {
	case *MultiError:
		httpError = ErrUnknownError
		var ss []string
		for i := range *e {
			ss = append(ss, ExtractHTTPError((*e)[i]).printLight())
		}
		httpError.Message = strings.Join(ss, ", ")
	case errorWithStack:
		httpError = e.httpError
	case Error:
		httpError = e
	default:
		httpError = ErrUnknownError
	}

	// if it's a custom err with no status use unknown error status
	if httpError.Status == 0 {
		httpError.Status = ErrUnknownError.Status
	}

	// if error's message is not empty do not override (custom message)
	// else set message for given accepted languages.
	if httpError.Message == "" {
		httpError.Message = httpError.Translate()
	}
	return httpError
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*MultiError); ok {
		err = NewError(ErrUnknownError, err)
	}

	// if it's already a Station does not override the error
	if e, ok := err.(errorWithStack); ok {
		return e
	}

	if e, ok := err.(Error); ok {
		return errorWithStack{
			root:      errors.New(e.Translate()),
			httpError: e,
		}
	}

	return errorWithStack{
		root:      errors.WithStack(err),
		httpError: ErrUnknownError,
	}
}

func (e *MultiError) Error() string {
	var ss []string
	for i := range *e {
		ss = append(ss, (*e)[i].Error())
	}
	return strings.Join(ss, ", ")
}

func Exit(format string, args ...interface{}) {
	if len(args) > 0 && format[:len(format)-1] != "\n" {
		format += "\n"
	}
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
