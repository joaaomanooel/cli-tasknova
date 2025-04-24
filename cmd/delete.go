package cmd

import (
	"fmt"

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
			tasks, err := task.Storage.Read(fileStorage)
			if err != nil {
				return fmt.Errorf("Error reading tasks: %e", err)
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
				return fmt.Errorf("Task with ID %d not found. \n", id)
			}

			if err := task.Storage.Save(fileStorage, newTasks); err != nil {
				return fmt.Errorf("Error saving tasks: %e", err)
			}

			fmt.Println("Task deleted successfully! 🎉")
			return nil
		},
	}

	cmd.Flags().UintP("id", "i", 0, "ID of the task to delete")
	cmd.MarkFlagRequired("id")

	return cmd
}

func init() {
	deleteTaskCmd()
	rootCmd.AddCommand(deleteTaskCmd())
}
