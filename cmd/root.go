package cmd

import (
	"fmt"
	"os"

	"github.com/joaaomanooel/cli-tasknova/internal/task"
	"github.com/spf13/cobra"
)

var (
	dataFile    = "tasks.json"
	fileStorage = &task.FileStorage{FilePath: dataFile}
)

var rootCmd = &cobra.Command{
	Use:   "tasknova",
	Short: "TaskNova CLI for managing tasks and notes",
	Long:  `TaskNova is a CLI tool for managing your tasks and notes with ease.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
