package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()

	// Test help command
	os.Args = []string{"tasknova", "--help"}
	assert.NotPanics(t, func() {
		Execute()
	})

	// Test version command
	os.Args = []string{"tasknova", "--version"}
	assert.NotPanics(t, func() {
		Execute()
	})

	// Test invalid command
	os.Args = []string{"tasknova", "--invalid-flag"}
	assert.NotPanics(t, func() {
		Execute()
	})
}
