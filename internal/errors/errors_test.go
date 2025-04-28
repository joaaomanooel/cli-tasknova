package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskError_Error(t *testing.T) {
	testCases := []struct {
		name     string
		taskErr  *TaskError
		expected string
	}{
		{
			name: "with underlying error",
			taskErr: &TaskError{
				Code:    "TEST_ERROR",
				Message: "test message",
				Err:     errors.New("underlying error"),
			},
			expected: "TEST_ERROR: test message (underlying error)",
		},
		{
			name: "without underlying error",
			taskErr: &TaskError{
				Code:    "TEST_ERROR",
				Message: "test message",
				Err:     nil,
			},
			expected: "TEST_ERROR: test message",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.taskErr.Error()

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestNewTaskError(t *testing.T) {
	code := "TEST_ERROR"
	message := "test message"
	err := errors.New("underlying error")

	taskErr := NewTaskError(code, message, err)

	assert.Equal(t, code, taskErr.Code)
	assert.Equal(t, message, taskErr.Message)
	assert.Equal(t, err, taskErr.Err)
}
