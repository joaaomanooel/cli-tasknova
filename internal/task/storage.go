package task

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/joaaomanooel/cli-tasknova/internal/constants"
	"github.com/joaaomanooel/cli-tasknova/internal/errors"
)

type FileStorage struct {
	FilePath string
	tasks    map[uint]Task
	mu       sync.RWMutex
}

type Storage interface {
	Save(tasks []Task) error
	Read() ([]Task, error)
	GetByID(id uint) (*Task, error)
	Update(task *Task) error
}

func (s *FileStorage) initTasksMap() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tasks != nil {
		return nil
	}

	tasks, err := s.Read()
	if err != nil {
		return err
	}

	s.tasks = make(map[uint]Task, len(tasks))
	for _, t := range tasks {
		s.tasks[t.ID] = t
	}
	return nil
}

func (s *FileStorage) GetByID(id uint) (*Task, error) {
	if err := s.initTasksMap(); err != nil {
		return nil, err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.NewTaskError(constants.NotFoundError, "Task not found", nil)
	}
	return &task, nil
}

func (s *FileStorage) Update(task *Task) error {
	if err := s.initTasksMap(); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = *task

	tasks := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}

	return s.Save(tasks)
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
