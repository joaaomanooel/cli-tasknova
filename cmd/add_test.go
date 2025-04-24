package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddCommandTestSuite struct {
	suite.Suite
	tempFile         string
	originalDataFile string
	buffer           *bytes.Buffer
}

func (s *AddCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}

	s.buffer = &bytes.Buffer{}

	rootCmd.ResetCommands()
	rootCmd.AddCommand(addTaskCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *AddCommandTestSuite) TearDownTest() {
	os.Remove(dataFile)
}

func (s *AddCommandTestSuite) TestAddTask() {
	// Arrange
	existingTasks := []task.Task{}
	err := task.Storage.Save(fileStorage, existingTasks)
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{
		"add",
		"--title", "Test Task",
		"--description", "Test Description",
		"--priority", "high",
	})

	err = rootCmd.Execute()
	assert.NoError(s.T(), err)

	tasks, err := task.Storage.Read(fileStorage)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), tasks, 1)

	newTask := tasks[0]
	assert.Equal(s.T(), "Test Task", newTask.Title)
	assert.Equal(s.T(), "Test Description", newTask.Description)
	assert.Equal(s.T(), "high", newTask.Priority)
	assert.NotZero(s.T(), newTask.ID) // Just verify ID exists
}

func (s *AddCommandTestSuite) TestAddTaskWithInvalidFile() {
	dataFile = "/nonexistent/directory/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}

	rootCmd.SetArgs([]string{
		"add",
		"--title", "Test Task",
		"--description", "Test Description",
		"--priority", "high",
	})

	err := rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "Error saving the task")
}

func (s *AddCommandTestSuite) TestAddTaskWithMissingTitle() {
	rootCmd.SetArgs([]string{
		"add",
		"--description", "Test Description",
		"--priority", "high",
	})

	err := rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "required flag(s) \"title\" not set")
}

func TestAddCommandSuite(t *testing.T) {
	suite.Run(t, new(AddCommandTestSuite))
}
