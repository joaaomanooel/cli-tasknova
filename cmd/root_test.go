package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RootCommandTestSuite struct {
	suite.Suite
	buffer          *bytes.Buffer
	tempFileStorage *task.FileStorage
}

func (s *RootCommandTestSuite) SetupTest() {
	s.tempFileStorage = fileStorage
	dataFile = testDataFile
	s.tempFileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = s.tempFileStorage

	s.buffer = &bytes.Buffer{}

	rootCmd.ResetCommands()
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *RootCommandTestSuite) TearDownTest() {
	if err := os.Remove(dataFile); err != nil {
		fmt.Printf("Failed to remove test file: %v\n", err)
	}
}

func TestMain(m *testing.M) {
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	os.Exit(m.Run())
}

func TestExecute(t *testing.T) {
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	err := rootCmd.Execute()

	assert.NoError(t, err, "Root command should execute without error")
}

func TestRootCommandSuite(t *testing.T) {
	suite.Run(t, new(RootCommandTestSuite))
}
