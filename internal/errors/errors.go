package errors

import "fmt"

type TaskError struct {
	Code    string
	Message string
	Err     error
}

func (e *TaskError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewTaskError(code, message string, err error) *TaskError {
	return &TaskError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
