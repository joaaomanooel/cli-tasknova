package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CompletionCommandTestSuite struct {
	suite.Suite
	buffer *bytes.Buffer
}

func (s *CompletionCommandTestSuite) SetupTest() {
	rootCmd = &cobra.Command{
		Use:   "tasknova",
		Short: "A CLI task manager",
	}

	s.buffer = &bytes.Buffer{}

	rootCmd.ResetCommands()
	rootCmd.AddCommand(completionCmd())
	rootCmd.SetOut(s.buffer)
	rootCmd.SetErr(s.buffer)
}

func (s *CompletionCommandTestSuite) TestCompletionNoArgs() {
	rootCmd.SetArgs([]string{})

	err := rootCmd.Execute()

	assert.NoError(s.T(), err)
	assert.Contains(s.T(), s.buffer.String(), "completion")
}

func (s *CompletionCommandTestSuite) TestCompletionBash() {
	rootCmd.SetArgs([]string{"completion", "bash"})

	err := rootCmd.Execute()

	assert.NoError(s.T(), err)
}

func (s *CompletionCommandTestSuite) TestCompletionZsh() {
	rootCmd.SetArgs([]string{"completion", "zsh"})

	err := rootCmd.Execute()

	assert.NoError(s.T(), err)
}

func TestCompletionCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CompletionCommandTestSuite))
}
