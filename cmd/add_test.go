package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddCommandTestSuite struct {
	suite.Suite
	buffer          *bytes.Buffer
	tempFileStorage *task.FileStorage
}

func (s *AddCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.tempFileStorage = fileStorage
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	s.buffer = &bytes.Buffer{}

	rootCmd.ResetCommands()
	rootCmd.AddCommand(addTaskCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *AddCommandTestSuite) TearDownTest() {
	if err := os.Remove(dataFile); err != nil {
		fmt.Printf("Failed to remove test file: %v\n", err)
	}
}

func (s *AddCommandTestSuite) TestAddTask() {

	existingTasks := []task.Task{}
	err := task.DefaultStorage.Save(existingTasks) // Use DefaultStorage
	assert.NoError(s.T(), err)

	rootCmd.SetArgs([]string{
		"add",
		"--title", "Test Task",
		"--description", "Test Description",
		"--priority", "high",
	})

	err = rootCmd.Execute()
	assert.NoError(s.T(), err)

	tasks, err := task.DefaultStorage.Read() // Use DefaultStorage
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), tasks, "Tasks should not be empty after adding")

	newTask := tasks[0]
	assert.Equal(s.T(), "Test Task", newTask.Title)
	assert.Equal(s.T(), "Test Description", newTask.Description)
	assert.Equal(s.T(), "high", newTask.Priority)
	assert.NotZero(s.T(), newTask.ID)
}

func (s *AddCommandTestSuite) TestAddTaskWithInvalidFile() {
	dataFile = "./add/invalid/path/tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	rootCmd.SetArgs([]string{
		"add",
		"--title", "Test Task",
		"--description", "Test Description",
		"--priority", "high",
	})

	err := rootCmd.Execute()

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "WRITE_ERROR: Directory does not exist")
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

func (s *AddCommandTestSuite) TestAddTaskWithEmptyTitle() {
	cmd := addTaskCmd()
	cmd.SetArgs([]string{"--title", ""})

	err := cmd.Execute()
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "title is required")
}

func TestAddCommandSuite(t *testing.T) {
	suite.Run(t, new(AddCommandTestSuite))
}
