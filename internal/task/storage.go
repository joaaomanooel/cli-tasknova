package task

import (
	"encoding/json"
	"os"

	"github.com/joaaomanooel/cli-tasknova/internal/errors"
	"github.com/joaaomanooel/cli-tasknova/internal/storage"
)

type FileStorage struct {
	FilePath string
}

type Storage interface {
	Save(tasks []Task) error
	Read() ([]Task, error)
}

func (fs *FileStorage) Save(tasks []Task) error {
	// Ensure storage directory exists
	if err := storage.EnsureStorageDirectory(fs.FilePath); err != nil {
		return errors.NewTaskError("STORAGE_DIR_ERROR", "Failed to create storage directory", err)
	}

	// Marshal tasks with indentation for better readability
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return errors.NewTaskError("MARSHAL_ERROR", "Failed to marshal tasks", err)
	}

	// Write file with proper permissions
	err = os.WriteFile(fs.FilePath, data, storage.DefaultFileMode)
	if err != nil {
		return errors.NewTaskError("WRITE_ERROR", "Failed to write tasks file", err)
	}

	return storage.EnsureFilePermissions(fs.FilePath)
}

func (fs *FileStorage) Read() ([]Task, error) {
	data, err := os.ReadFile(fs.FilePath)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, errors.NewTaskError("READ_ERROR", "Failed to read tasks file", err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, errors.NewTaskError("UNMARSHAL_ERROR", "Failed to parse tasks data", err)
	}

	return tasks, nil
}
