package task

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileStorageTestSuite struct {
	suite.Suite
	storage *FileStorage
	tmpFile string
}

func (s *FileStorageTestSuite) SetupTest() {
	s.tmpFile = "test_tasks.json"
	s.storage = &FileStorage{FilePath: s.tmpFile}
}

func (s *FileStorageTestSuite) TearDownTest() {
	os.Remove(s.tmpFile)
}

func (s *FileStorageTestSuite) TestSaveAndRead() {
	expectedTask := Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks := []Task{expectedTask}

	saveErr := s.storage.Save(tasks)
	readTasks, readErr := s.storage.Read()

	assert.NoError(s.T(), saveErr)
	assert.NoError(s.T(), readErr)
	assert.Len(s.T(), readTasks, 1)
	actualTask := readTasks[0]
	assert.Equal(s.T(), expectedTask.ID, actualTask.ID)
	assert.Equal(s.T(), expectedTask.Title, actualTask.Title)
	assert.Equal(s.T(), expectedTask.Description, actualTask.Description)
	assert.Equal(s.T(), expectedTask.Priority, actualTask.Priority)
}

func (s *FileStorageTestSuite) TestReadEmptyFile() {
	tasks, err := s.storage.Read()

	assert.NoError(s.T(), err)
	assert.Empty(s.T(), tasks)
}

func (s *FileStorageTestSuite) TestSaveInvalidPath() {
	s.storage.FilePath = "/invalid/path/tasks.json"
	tasks := []Task{{ID: 1, Title: "Test"}}

	err := s.storage.Save(tasks)

	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "STORAGE_DIR_ERROR: Failed to create storage directory")
}

func (s *FileStorageTestSuite) TestReadInvalidJSON() {
	err := os.WriteFile(s.tmpFile, []byte("invalid json content"), 0644)
	assert.NoError(s.T(), err)

	tasks, err := s.storage.Read()

	assert.Error(s.T(), err)
	assert.Empty(s.T(), tasks)
}

func TestFileStorageSuite(t *testing.T) {
	suite.Run(t, new(FileStorageTestSuite))
}
