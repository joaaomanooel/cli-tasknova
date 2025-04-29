package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateCommandTestSuite struct {
	suite.Suite
	buffer          *bytes.Buffer
	tempFileStorage *task.FileStorage
}

func (s *UpdateCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.tempFileStorage = fileStorage
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	rootCmd.ResetCommands()
	rootCmd.AddCommand(updateTaskCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)

	s.buffer = &bytes.Buffer{}
}

func (s *UpdateCommandTestSuite) TearDownTest() {
	if err := os.Remove(dataFile); err != nil {
		fmt.Printf("Failed to remove temporary file: %v\n", err)
	}
}

func (s *UpdateCommandTestSuite) TestUpdateTask() {
	existingTask := task.Task{
		ID:          1234567891,
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := task.DefaultStorage.Save([]task.Task{existingTask})
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{
		"update",
		"--id", "1234567891",
		"--title", "Updated Task",
		"--description", "Updated Description",
		"--priority", "low",
	})

	err = rootCmd.Execute()
	assert.NoError(s.T(), err)

	tasks, _ := task.DefaultStorage.Read()

	updatedTask := tasks[0]
	assert.Equal(s.T(), "Updated Task", updatedTask.Title)
	assert.Equal(s.T(), "Updated Description", updatedTask.Description)
	assert.Equal(s.T(), "low", updatedTask.Priority)
	assert.Equal(s.T(), existingTask.ID, updatedTask.ID)
}

func (s *UpdateCommandTestSuite) TestUpdateTaskWithInvalidID() {
	rootCmd.SetArgs([]string{
		"update",
		"--id", "987654321",
		"--title", "Updated Task",
	})
	err := rootCmd.Execute()
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "Task not found")
}

func (s *UpdateCommandTestSuite) TestUpdateTaskWithEmptyFields() {
	existingTask := task.Task{
		ID:          1234567891,
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := task.DefaultStorage.Save([]task.Task{existingTask})
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{
		"update",
		"--id", "1234567891",
		"--title", "",
		"--description", "",
		"--priority", "",
	})
	err = rootCmd.Execute()
	assert.NoError(s.T(), err)

	tasks, err := task.DefaultStorage.Read()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), tasks, 1)

	updatedTask := tasks[0]
	assert.Equal(s.T(), "Test Task", updatedTask.Title)
	assert.Equal(s.T(), "Test Description", updatedTask.Description)
	assert.Equal(s.T(), "high", updatedTask.Priority)
	assert.Equal(s.T(), existingTask.ID, updatedTask.ID)
}

func (s *UpdateCommandTestSuite) TestUpdateTaskWithoutId() {
	rootCmd.SetArgs([]string{
		"update",
		"--title", "Updated Task",
	})
	err := rootCmd.Execute()
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "required flag(s) \"id\" not set")
}

func (s *UpdateCommandTestSuite) TestUpdateTaskWithInvalidFile() {
	dataFile = "./update/invalid/path/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	rootCmd.SetArgs([]string{
		"update",
		"--id", "1234567891",
		"--title", "Updated Task",
	})

	err := rootCmd.Execute()
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "Directory does not exist")
}

func TestUpdateCommandTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateCommandTestSuite))
}
