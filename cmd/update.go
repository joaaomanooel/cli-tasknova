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

			tasks, err := task.DefaultStorage.Read()
			if err != nil {
				return err
			}

			updated := false
			for i, t := range tasks {
				if t.ID == id {
					if title != "" {
						t.Title = title
					}
					if description != "" {
						t.Description = description
					}
					if priority != "" {
						t.Priority = priority
					}

					t.UpdatedAt = time.Now()
					tasks[i] = t
					updated = true

					break
				}
			}

			if !updated {
				return errors.NewTaskError(constants.NotFoundError, "Task not found", err)
			}

			if err := task.DefaultStorage.Save(tasks); err != nil {
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
