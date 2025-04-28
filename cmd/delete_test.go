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

type DeleteCommandTestSuite struct {
	suite.Suite
	buffer *bytes.Buffer
}

func (s *DeleteCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	s.buffer = &bytes.Buffer{}

	rootCmd.ResetCommands()
	rootCmd.AddCommand(deleteTaskCmd()) // Changed from addTaskCmd() to deleteTaskCmd()
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *DeleteCommandTestSuite) TearDownTest() {
	os.Remove(dataFile)
}

func (s *DeleteCommandTestSuite) TestDeleteTask() {
	tasks := []task.Task{
		{ID: 1, Title: "Test Task", Description: "Test Description", Priority: "high", CreatedAt: time.Now()},
	}
	err := task.DefaultStorage.Save(tasks)
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{"delete", "--id", "1"})

	err = rootCmd.Execute()

	assert.NoError(s.T(), err)
	remaining, err := task.DefaultStorage.Read()
	assert.NoError(s.T(), err)
	assert.Empty(s.T(), remaining, "Tasks should be empty after deletion")
}

func (s *DeleteCommandTestSuite) TestDeleteTaskWithInvalidFile() {
	dataFile = "/dev/null/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	rootCmd.SetArgs([]string{"delete", "--id", "1"})

	err := rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "READ_ERROR: Failed to read tasks file")
}

func (s *DeleteCommandTestSuite) TestDeleteTaskWithInvalidID() {
	tasks := []task.Task{
		{ID: 1, Title: "Test Task", Description: "Test Description", Priority: "high", CreatedAt: time.Now()},
	}
	err := task.DefaultStorage.Save(tasks)
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{"delete", "--id", "invalid"})

	err = rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "invalid argument \"invalid\" for \"-i, --id\" flag: strconv.ParseUint")
}

func (s *DeleteCommandTestSuite) TestDeleteNonExistentTask() {
	tasks := []task.Task{
		{ID: 1, Title: "Test Task", Description: "Test Description", Priority: "high", CreatedAt: time.Now()},
	}
	err := task.DefaultStorage.Save(tasks)
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{"delete", "--id", "999"})

	err = rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "Task with ID 999 not found")
}

func TestDeleteCommandSuite(t *testing.T) {
	suite.Run(t, new(DeleteCommandTestSuite))
}
