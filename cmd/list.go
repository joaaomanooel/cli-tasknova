package cmd

import (
	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

func listTasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks and notes",
		Long:  `Displays all tasks and notes stored in tasks.json file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := task.DefaultStorage.Read()
			if err != nil {
				return errors.NewTaskError(constants.ReadError, "Failed to read tasks", err)
			}

			if len(tasks) == 0 {
				cmd.Println("No tasks found")
				return nil
			}

			cmd.Println("Your Tasks:")
			for _, task := range tasks {
				cmd.Printf("ID: %d\n", task.ID)
				cmd.Printf("Title: %s\n", task.Title)
				cmd.Printf("Description: %s\n", task.Description)
				cmd.Printf("Priority: %s\n", task.Priority)
				cmd.Printf("Created At: %s\n", task.CreatedAt.Format("Mon, 02 Jan 2006 15:04"))
				cmd.Printf("Updated At: %s\n", task.UpdatedAt.Format("Mon, 02 Jan 2006 15:04"))
				cmd.Println("------------------------")
			}
			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(listTasksCmd())
}
