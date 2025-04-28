package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskTestSuite struct {
	suite.Suite
	now    time.Time
	future time.Time
	past   time.Time
}

func (s *TaskTestSuite) SetupTest() {
	s.now = time.Now()
	s.future = s.now.Add(time.Hour)
	s.past = s.now.Add(-time.Hour)
}

func (s *TaskTestSuite) TestValidTaskWithAllFields() {
	task := Task{
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		CreatedAt:   s.now,
		UpdatedAt:   s.now,
	}

	err := task.Validate()
	assert.NoError(s.T(), err)
}

func (s *TaskTestSuite) TestEmptyTitle() {

	task := Task{
		Priority:  "high",
		CreatedAt: s.now,
		UpdatedAt: s.now,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "title is required")
}

func (s *TaskTestSuite) TestInvalidPriority() {

	task := Task{
		Title:     "Test Task",
		Priority:  "invalid",
		CreatedAt: s.now,
		UpdatedAt: s.now,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "invalid priority")
}

func (s *TaskTestSuite) TestMissingCreatedAt() {

	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		UpdatedAt: s.now,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "created at is required")
}

func (s *TaskTestSuite) TestMissingUpdatedAt() {

	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		CreatedAt: s.now,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "updated at is required")
}

func (s *TaskTestSuite) TestValidPriorities() {
	validPriorities := []string{"low", "medium", "high"}

	for _, priority := range validPriorities {
		s.Run(priority, func() {

			task := Task{
				Title:     "Test Task",
				Priority:  priority,
				CreatedAt: s.now,
				UpdatedAt: s.now,
			}

			err := task.Validate()

			assert.NoError(s.T(), err)
		})
	}
}

func (s *TaskTestSuite) TestValidTimeSequence() {

	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		CreatedAt: s.past,
		UpdatedAt: s.now,
	}

	err := task.Validate()

	assert.NoError(s.T(), err)
}

func (s *TaskTestSuite) TestInvalidTimeSequence() {

	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		CreatedAt: s.now,
		UpdatedAt: s.past,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
}

func (s *TaskTestSuite) TestFutureCreationTime() {

	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		CreatedAt: s.future,
		UpdatedAt: s.future,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
}

func (s *TaskTestSuite) TestValidateFutureUpdatedAt() {

	future := time.Now().Add(time.Hour)
	task := Task{
		Title:     "Test Task",
		Priority:  "high",
		CreatedAt: time.Now(),
		UpdatedAt: future,
	}

	err := task.Validate()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "updated at cannot be in the future")
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}
