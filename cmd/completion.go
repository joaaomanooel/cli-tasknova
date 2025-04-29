package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func completionCmd() *cobra.Command {
	validShells := []string{"bash", "zsh"}
	example := `
  # Bash completion
  source <(tasknova completion bash)

  # Zsh completion
  source <(tasknova completion zsh)`

	cmd := &cobra.Command{
		Use:       "completion [bash|zsh]",
		Short:     "Generate completion script",
		Long:      `Generate shell completion script for TaskNova CLI.`,
		ValidArgs: validShells,
		Example:   example,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("shell type required. Valid shells: %v", validShells)
			}

			shell := args[0]
			switch shell {
			case "bash":
				return rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				return rootCmd.GenZshCompletion(os.Stdout)
			default:
				return fmt.Errorf("invalid shell type %q. Valid shells: %v", shell, validShells)
			}
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(completionCmd())
}
