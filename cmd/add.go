package cmd

import (
	"fmt"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

var (
	dataFile    = "tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
)

func addTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task or note",
		Long:  `Adds a new task by saving title, description, priority, and creation date.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			priority, _ := cmd.Flags().GetString("priority")

			if title == "" {
				return errors.NewTaskError(constants.ValidationError, "title is required", nil)
			}

			tasks, err := task.DefaultStorage.Read()
			if err != nil {
				return errors.NewTaskError(constants.ReadError, "Failed to read tasks", err)
			}

			newTask := task.Task{
				ID:          task.DefaultIDGenerator.GenerateID(),
				Title:       title,
				Description: description,
				Priority:    priority,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			tasks = append(tasks, newTask)
			if err := task.DefaultStorage.Save(tasks); err != nil {
				return errors.NewTaskError(constants.SaveError, "Failed to save task", err)
			}

			fmt.Println("Task added successfully! 🎉")
			return nil
		},
	}

	cmd.Flags().StringP("title", "t", "", "Title of the task")
	cmd.Flags().StringP("description", "d", "", "Description of the task")
	cmd.Flags().StringP("priority", "p", "low", "Priority of the task (low, normal, high)")
	if err := cmd.MarkFlagRequired("title"); err != nil {
		panic(fmt.Sprintf("Failed to mark title flag as required: %v", err))
	}

	return cmd
}

func init() {
	task.DefaultStorage = fileStorage
	rootCmd.AddCommand(addTaskCmd())
}
