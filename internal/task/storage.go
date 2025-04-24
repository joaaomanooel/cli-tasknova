package task

import (
	"encoding/json"
	"os"
)

type Storage interface {
	Save(tasks []Task) error
	Read() ([]Task, error)
}

type FileStorage struct {
	FilePath string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{FilePath: path}
}

func (fs *FileStorage) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, data, 0644)
}

func (fs *FileStorage) Read() ([]Task, error) {
	if _, err := os.Stat(fs.FilePath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
