package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskErrorTestSuite struct {
	suite.Suite
	baseError error
}

func (s *TaskErrorTestSuite) SetupTest() {
	s.baseError = errors.New("underlying error")
}

func (s *TaskErrorTestSuite) TestErrorWithUnderlyingError() {

	taskErr := &TaskError{
		Code:    "TEST_ERROR",
		Message: "test message",
		Err:     s.baseError,
	}
	expected := "TEST_ERROR: test message (underlying error)"

	result := taskErr.Error()

	assert.Equal(s.T(), expected, result)
}

func (s *TaskErrorTestSuite) TestErrorWithoutUnderlyingError() {

	taskErr := &TaskError{
		Code:    "TEST_ERROR",
		Message: "test message",
		Err:     nil,
	}
	expected := "TEST_ERROR: test message"

	result := taskErr.Error()

	assert.Equal(s.T(), expected, result)
}

func (s *TaskErrorTestSuite) TestNewTaskErrorCreation() {

	code := "TEST_ERROR"
	message := "test message"

	taskErr := NewTaskError(code, message, s.baseError)

	assert.Equal(s.T(), code, taskErr.Code)
	assert.Equal(s.T(), message, taskErr.Message)
	assert.Equal(s.T(), s.baseError, taskErr.Err)
}

func (s *TaskErrorTestSuite) TestNewTaskErrorWithNilError() {

	code := "TEST_ERROR"
	message := "test message"

	taskErr := NewTaskError(code, message, nil)

	assert.Equal(s.T(), code, taskErr.Code)
	assert.Equal(s.T(), message, taskErr.Message)
	assert.Nil(s.T(), taskErr.Err)
}

func TestTaskErrorSuite(t *testing.T) {
	suite.Run(t, new(TaskErrorTestSuite))
}
