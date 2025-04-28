package task

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
)

type FileStorage struct {
	FilePath string
}

type Storage interface {
	Save(tasks []Task) error
	Read() ([]Task, error)
}

func (fs *FileStorage) directoryExists() error {
	dir := filepath.Dir(fs.FilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return errors.NewTaskError(constants.WriteError, "Directory does not exist", err)
	}

	return nil
}

func (fs *FileStorage) Save(tasks []Task) error {
	if err := fs.directoryExists(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return errors.NewTaskError(constants.MarshalError, "Failed to marshal tasks", err)
	}

	// Use more restrictive file permissions (0600)
	if err := os.WriteFile(fs.FilePath, data, 0600); err != nil {
		return errors.NewTaskError(constants.WriteError, "Failed to write tasks file", err)
	}

	return nil
}

func (fs *FileStorage) Read() ([]Task, error) {
	if err := fs.directoryExists(); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(fs.FilePath)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, errors.NewTaskError(constants.ReadError, "Failed to read tasks file", err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, errors.NewTaskError(constants.UnmarshalError, "Failed to parse tasks data", err)
	}

	return tasks, nil
}
