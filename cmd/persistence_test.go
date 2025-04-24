package cmd

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PersistenceTestSuite struct {
	suite.Suite
	tempFile         string
	originalDataFile string
}

func (s *PersistenceTestSuite) SetupTest() {
	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
}

func (s *PersistenceTestSuite) TearDownTest() {
	os.Remove(dataFile)
}

func (s *PersistenceTestSuite) TestSaveAndReadTasks() {
	expectedTasks := []task.Task{
		{
			ID:          1,
			Title:       "Test Task",
			Description: "Test read and write functions",
			Priority:    "low",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	err := task.Storage.Save(fileStorage, expectedTasks)
	assert.NoError(s.T(), err, "Should save tasks without error")

	actualTasks, err := task.Storage.Read(fileStorage)

	assert.NoError(s.T(), err, "Should read tasks without error")

	expected, err := json.MarshalIndent(expectedTasks, "", "")
	assert.NoError(s.T(), err, "Should marshal expected tasks")

	actual, err := json.MarshalIndent(actualTasks, "", "")
	assert.NoError(s.T(), err, "Should marshal actual tasks")

	assert.Equal(s.T(), string(expected), string(actual), "Tasks read should match tasks written")
}

func (s *PersistenceTestSuite) TestReadTasksFromNonExistentFile() {
	dataFile = "non_existent_file.json"

	tasks, err := task.Storage.Read(fileStorage)

	assert.NoError(s.T(), err, "Reading non-existent file should return empty tasks without error")
	assert.Empty(s.T(), tasks, "Reading non-existent file should return empty tasks slice")
}

func (s *PersistenceTestSuite) TestSaveTasksToInvalidPath() {
	dataFile = "/invalid/path/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}

	tasks := []task.Task{{
		ID:          1,
		Title:       "Test Task",
		Description: "Test invalid path",
		Priority:    "low",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}}

	err := task.Storage.Save(fileStorage, tasks)

	assert.Error(s.T(), err, "Saving to invalid path should return error")
}

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(PersistenceTestSuite))
}
