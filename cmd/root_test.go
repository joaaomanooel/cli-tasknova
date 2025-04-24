package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	defer func() {
		if err := os.Remove(dataFile); err != nil {
			fmt.Printf("Failed to remove test file: %v\n", err)
		}
	}()

	os.Exit(m.Run())
}

func TestExecute(t *testing.T) {
	dataFile = testDataFile
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	defer func() {
		if err := os.Remove(dataFile); err != nil {
			fmt.Printf("Failed to remove test file: %v\n", err)
		}
	}()

	err := rootCmd.Execute()

	assert.NoError(t, err, "Root command should execute without error")
}
