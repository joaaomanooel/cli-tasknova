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

			if err := task.DefaultStorage.Delete(id); err != nil {
				return err
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
