package cmd

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListCommandTestSuite struct {
	suite.Suite
	tempFile         string
	originalDataFile string
	buffer           *bytes.Buffer
}

func (s *ListCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.buffer = &bytes.Buffer{}

	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}

	// Reset and setup root command
	rootCmd.ResetCommands()
	rootCmd.AddCommand(listTasksCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
	rootCmd.SetArgs([]string{"list"})
}

func (s *ListCommandTestSuite) TearDownTest() {
	os.Remove(dataFile)
}

func (s *ListCommandTestSuite) TestListSingleTask() {
	// Arrange
	newTask := task.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    "High",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := task.Storage.Save(fileStorage, []task.Task{newTask})
	assert.NoError(s.T(), err, "Failed to save tasks")

	// Act
	err = rootCmd.Execute()

	// Assert
	assert.NoError(s.T(), err)
	output := s.buffer.String()

	expectedStrings := []string{
		"ID: 1",
		"Title: Test Task",
		"Description: This is a test task",
		"Priority: High",
	}

	for _, expected := range expectedStrings {
		assert.Contains(s.T(), output, expected)
	}
}

func (s *ListCommandTestSuite) TestListEmptyTaskList() {
	// Arrange
	err := task.Storage.Save(fileStorage, []task.Task{})
	assert.NoError(s.T(), err, "Failed to save empty task list")

	// Act
	err = rootCmd.Execute()

	// Assert
	assert.NoError(s.T(), err)
	assert.Contains(s.T(), s.buffer.String(), "No tasks found")
}

func (s *ListCommandTestSuite) TestListMultipleTasks() {
	// Arrange
	tasks := []task.Task{
		{
			ID:          1,
			Title:       "First Task",
			Description: "This is the first task",
			Priority:    "High",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Second Task",
			Description: "This is the second task",
			Priority:    "Low",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	err := task.Storage.Save(fileStorage, tasks)
	assert.NoError(s.T(), err, "Failed to save tasks")

	// Act
	err = rootCmd.Execute()

	// Assert
	assert.NoError(s.T(), err)
	output := s.buffer.String()

	expectedStrings := []string{
		"ID: 1",
		"Title: First Task",
		"Description: This is the first task",
		"Priority: High",
		"ID: 2",
		"Title: Second Task",
		"Description: This is the second task",
		"Priority: Low",
	}

	for _, expected := range expectedStrings {
		assert.Contains(s.T(), output, expected)
	}
}

func (s *ListCommandTestSuite) TestListTasksWithInvalidFile() {
	// Arrange
	dataFile = "/dev/null/tasks.json" // This path will always be invalid on Unix systems
	fileStorage = &task.FileStorage{FilePath: dataFile}

	// Create the invalid directory to ensure it exists but is not writable
	err := os.MkdirAll("/dev/null", 0400)
	if err == nil {
		defer os.RemoveAll("/dev/null")
	}

	// Act
	err = rootCmd.Execute()

	// Assert
	assert.Error(s.T(), err, "Expected error when reading from invalid file path")
	assert.Contains(s.T(), err.Error(), "error reading tasks", "Error message should indicate task reading failure")
}

func TestListCommandSuite(t *testing.T) {
	suite.Run(t, new(ListCommandTestSuite))
}
