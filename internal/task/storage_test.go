package task

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileStorageTestSuite struct {
	suite.Suite
	storage  *FileStorage
	tmpFile  string
	testTask Task
}

func (s *FileStorageTestSuite) SetupTest() {

	s.tmpFile = "test_tasks.json"
	s.storage = &FileStorage{FilePath: s.tmpFile}
	s.testTask = Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (s *FileStorageTestSuite) TearDownTest() {
	if err := os.Remove(s.tmpFile); err != nil {
		fmt.Printf("Failed to remove test file: %v\n", err)
	}
}

func (s *FileStorageTestSuite) TestSaveAndRead() {

	tasks := []Task{s.testTask}

	saveErr := s.storage.Save(tasks)
	readTasks, readErr := s.storage.Read()

	assert.NoError(s.T(), saveErr)
	assert.NoError(s.T(), readErr)
	assert.Len(s.T(), readTasks, 1)

	actualTask := readTasks[0]
	assert.Equal(s.T(), s.testTask.ID, actualTask.ID)
	assert.Equal(s.T(), s.testTask.Title, actualTask.Title)
	assert.Equal(s.T(), s.testTask.Description, actualTask.Description)
	assert.Equal(s.T(), s.testTask.Priority, actualTask.Priority)
}

func (s *FileStorageTestSuite) TestReadEmptyFile() {

	tasks, err := s.storage.Read()

	assert.NoError(s.T(), err)
	assert.Empty(s.T(), tasks)
}

func (s *FileStorageTestSuite) TestSaveInvalidPath() {

	invalidPath := filepath.Join("storage", "invalid", "path", "tasks.json")
	storage := &FileStorage{FilePath: invalidPath}

	err := storage.Save([]Task{s.testTask})

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "WRITE_ERROR: Directory does not exist")
}

func (s *FileStorageTestSuite) TestReadInvalidJSON() {

	err := os.WriteFile(s.tmpFile, []byte("invalid json content"), 0644)
	assert.NoError(s.T(), err)

	tasks, err := s.storage.Read()

	assert.Error(s.T(), err)
	assert.Empty(s.T(), tasks)
}

func (s *FileStorageTestSuite) TestReadWithPermissionError() {

	err := os.WriteFile(s.tmpFile, []byte("{}"), 0000)
	assert.NoError(s.T(), err)

	tasks, err := s.storage.Read()

	assert.Error(s.T(), err)
	assert.Nil(s.T(), tasks)
	assert.Contains(s.T(), err.Error(), "Failed to read tasks file")
}

func (s *FileStorageTestSuite) TestSaveWithWriteError() {
	err := os.MkdirAll(s.tmpFile, 0755)
	assert.NoError(s.T(), err)

	err = s.storage.Save([]Task{s.testTask})

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "Failed to write tasks file")
}

func TestFileStorageSuite(t *testing.T) {
	suite.Run(t, new(FileStorageTestSuite))
}
