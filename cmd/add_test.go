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
	tempFile            string
	originalDataFile    string
	buffer              *bytes.Buffer
	originalFileStorage *task.FileStorage
	tempFileStorage     *task.FileStorage
}

func (s *AddCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.tempFileStorage = fileStorage
	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage  // Add this line to initialize default storage

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
	err := task.DefaultStorage.Save(existingTasks)  // Use DefaultStorage
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{
		"add",
		"--title", "Test Task",
		"--description", "Test Description",
		"--priority", "high",
	})

	err = rootCmd.Execute()
	assert.NoError(s.T(), err)

	tasks, err := task.DefaultStorage.Read()  // Use DefaultStorage
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), tasks, "Tasks should not be empty after adding")

	newTask := tasks[0]
	assert.Equal(s.T(), "Test Task", newTask.Title)
	assert.Equal(s.T(), "Test Description", newTask.Description)
	assert.Equal(s.T(), "high", newTask.Priority)
	assert.NotZero(s.T(), newTask.ID)
}

func (s *AddCommandTestSuite) TestAddTaskWithInvalidFile() {
    // Arrange
    dataFile = "/dev/null/tasks.json"
    fileStorage = &task.FileStorage{FilePath: dataFile}
    task.DefaultStorage = fileStorage

    rootCmd.SetArgs([]string{
        "add",
        "--title", "Test Task",
        "--description", "Test Description",
        "--priority", "high",
    })

    // Act
    err := rootCmd.Execute()

    // Assert
    assert.Error(s.T(), err)
    assert.Contains(s.T(), err.Error(), "READ_ERROR: Failed to read tasks file")
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
