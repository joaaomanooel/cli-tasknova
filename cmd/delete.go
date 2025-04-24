package cmd

import (
	"fmt"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

func deleteTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task or note by ID",
		Long:  `Removes the task with the specified ID from the storage file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetUint("id")
			tasks, err := task.DefaultStorage.Read()
			if err != nil {
				return errors.NewTaskError(constants.ReadError, "Failed to read tasks", err)
			}

			newTasks := []task.Task{}
			found := false
			for _, t := range tasks {
				if t.ID == id {
					found = true
					continue
				}

				newTasks = append(newTasks, t)
			}

			if !found {
				return errors.NewTaskError(constants.NotFoundError, fmt.Sprintf("Task with ID %d not found", id), nil)
			}

			if err := task.DefaultStorage.Save(newTasks); err != nil {
				return errors.NewTaskError(constants.SaveError, "Failed to save tasks", err)
			}

			fmt.Println("Task deleted successfully! 🎉")
			return nil
		},
	}

	cmd.Flags().UintP("id", "i", 0, "ID of the task to delete")
	if err := cmd.MarkFlagRequired("id"); err != nil {
		panic(fmt.Sprintf("Failed to mark id flag as required: %v", err))
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd())
}
