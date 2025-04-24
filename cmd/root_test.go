package cmd

import (
	"os"
	"testing"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	dataFile = "temp_tasks_test.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
	task.DefaultStorage = fileStorage

	defer os.Remove(dataFile)

	err := rootCmd.Execute()

	assert.NoError(t, err, "Root command should execute without error")
}
