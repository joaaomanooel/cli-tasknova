package cmd

import (
	"fmt"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/olekukonko/tablewriter"
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

			table := tablewriter.NewWriter(cmd.OutOrStdout())
			table.SetHeader([]string{"ID", "Title", "Description", "Priority", "Created At", "Updated At"})
			table.SetRowLine(true)

			cmd.Println("Your Tasks:")
			for _, task := range tasks {
				table.Append([]string{
					fmt.Sprintf("%d", task.ID),
					task.Title,
					task.Description,
					task.Priority,
					task.CreatedAt.Format("Mon, 02 Jan 2006 15:04"),
					task.UpdatedAt.Format("Mon, 02 Jan 2006 15:04"),
				})
			}

			table.Render()
			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(listTasksCmd())
}
