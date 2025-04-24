package cmd

import (
	"fmt"
	"time"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

var dataFile = "tasks.json"
var fileStorage = &task.FileStorage{FilePath: dataFile}

func addTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task or note",
		Long:  `Adds a new task by saving title, description, priority, and creation date.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")
			priority, _ := cmd.Flags().GetString("priority")

			fmt.Println("Title:", title)
			fmt.Println("Description:", description)
			fmt.Println("Priority:", priority)
			fmt.Println("Created At:", time.Now().Format("2006-01-02 15:04:05"))

			if title == "" {
				return fmt.Errorf("title is required")
			}

			newTask := task.Task{
				ID:          generateID(),
				Title:       title,
				Description: description,
				Priority:    priority,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			tasks, err := task.Storage.Read(fileStorage)

			if err != nil {
				return fmt.Errorf("Error reading tasks file: %v", err)
			}

			tasks = append(tasks, newTask)
			err = task.Storage.Save(fileStorage, tasks)

			if err != nil {
				return fmt.Errorf("Error saving the task: %v", err)
			}

			fmt.Println("Task added successfully! 🎉")
			return nil
		},
	}

	cmd.Flags().StringP("title", "t", "", "Title of the task")
	cmd.Flags().StringP("description", "d", "", "Description of the task")
	cmd.Flags().StringP("priority", "p", "low", "Priority of the task (low, normal, high)")
	cmd.MarkFlagRequired("title")

	return cmd
}

func init() {
	addTaskCmd()
	rootCmd.AddCommand(addTaskCmd())
}

func generateID() uint {
	return uint(time.Now().UnixNano() / 1e6)
}
