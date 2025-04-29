package cmd

import (
	"fmt"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

func updateTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates an existing task or note",
		Long:  `Updates the title, description or priority of an existing task by its ID.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetUint("id")
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			priority, _ := cmd.Flags().GetString("priority")

			currentTask, err := task.DefaultStorage.GetByID(id)
			if err != nil {
				return err
			}

			if title != "" {
				currentTask.Title = title
			}
			if description != "" {
				currentTask.Description = description
			}
			if priority != "" {
				currentTask.Priority = priority
			}

			currentTask.UpdatedAt = time.Now()

			if err := task.DefaultStorage.Update(currentTask); err != nil {
				return errors.NewTaskError(constants.UpdateError, "Failed to update task", err)
			}

			fmt.Printf("Task %v updated successfully! 🎉", id)
			return nil
		},
	}

	cmd.Flags().UintP("id", "i", 0, "ID of the task to update")
	cmd.Flags().StringP("title", "t", "", "New title for the task")
	cmd.Flags().StringP("description", "d", "", "New description for the task")
	cmd.Flags().StringP("priority", "p", "", "New priority for the task")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		panic(fmt.Sprintln("Error marking flag as required:", err))
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(updateTaskCmd())
}
