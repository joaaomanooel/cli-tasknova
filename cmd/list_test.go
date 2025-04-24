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
	buffer *bytes.Buffer
}

func (s *ListCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.buffer = &bytes.Buffer{}

	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage // Add this line to initialize default storage

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

	err := task.DefaultStorage.Save(tasks) // Use DefaultStorage instead of Storage interface
	assert.NoError(s.T(), err, "Failed to save tasks")

	// Act
	err = rootCmd.Execute()

	// Assert
	assert.NoError(s.T(), err)
	output := s.buffer.String()

	expectedStrings := []string{
		"Title: First Task",
		"Description: This is the first task",
		"Priority: High",
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
	dataFile = "/dev/null/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	// Act
	err := rootCmd.Execute()

	// Assert
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "READ_ERROR: Failed to read tasks file")
}

func TestListCommandSuite(t *testing.T) {
	suite.Run(t, new(ListCommandTestSuite))
}
