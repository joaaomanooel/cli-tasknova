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

	dataFile = testDataFile
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
	if err := os.Remove(dataFile); err != nil {
		fmt.Printf("Failed to remove test file: %v\n", err)
	}
}

func (s *ListCommandTestSuite) TestListSingleTask() {
	now := time.Now()
	newTask := task.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    "High",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := task.DefaultStorage.Save([]task.Task{newTask})
	assert.NoError(s.T(), err, "Failed to save tasks")

	err = rootCmd.Execute()

	assert.NoError(s.T(), err)
	output := s.buffer.String()

	assert.Contains(s.T(), output, "Your Tasks:")
	assert.Contains(s.T(), output, "1")
	assert.Contains(s.T(), output, "Test Task")
	assert.Contains(s.T(), output, "This is a test task")
	assert.Contains(s.T(), output, "High")
	assert.Contains(s.T(), output, now.Format("Mon, 02 Jan 2006 15:04"))
}

func (s *ListCommandTestSuite) TestListEmptyTaskList() {
	err := task.DefaultStorage.Save([]task.Task{})
	assert.NoError(s.T(), err, "Failed to save empty task list")

	err = rootCmd.Execute()

	assert.NoError(s.T(), err)
	assert.Contains(s.T(), s.buffer.String(), "No tasks found")
}

func (s *ListCommandTestSuite) TestListMultipleTasks() {
	now := time.Now()
	tasks := []task.Task{
		{
			ID:          1,
			Title:       "First Task",
			Description: "This is the first task",
			Priority:    "High",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Title:       "Second Task",
			Description: "This is the second task",
			Priority:    "Low",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	err := task.DefaultStorage.Save(tasks)
	assert.NoError(s.T(), err, "Failed to save tasks")

	err = rootCmd.Execute()

	assert.NoError(s.T(), err)
	output := s.buffer.String()

	assert.Contains(s.T(), output, "Your Tasks:")
	assert.Contains(s.T(), output, "1")
	assert.Contains(s.T(), output, "First Task")
	assert.Contains(s.T(), output, "This is the first task")
	assert.Contains(s.T(), output, "High")
	assert.Contains(s.T(), output, "2")
	assert.Contains(s.T(), output, "Second Task")
	assert.Contains(s.T(), output, "This is the second task")
	assert.Contains(s.T(), output, "Low")
	assert.Contains(s.T(), output, now.Format("Mon, 02 Jan 2006 15:04"))
}

func (s *ListCommandTestSuite) TestListTasksWithInvalidFile() {
	invalidPath := "./list/invalid/path/tasks.json"
	dataFile = invalidPath
	fileStorage = &task.FileStorage{FilePath: invalidPath}
	task.DefaultStorage = fileStorage

	rootCmd.SetArgs([]string{"list"})
	err := rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "WRITE_ERROR: Directory does not exist")
}

func TestListCommandSuite(t *testing.T) {
	suite.Run(t, new(ListCommandTestSuite))
}
