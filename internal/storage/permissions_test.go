package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PermissionsTestSuite struct {
	suite.Suite
	tempDir string
}

func (s *PermissionsTestSuite) SetupTest() {
	s.tempDir = filepath.Join(os.TempDir(), "tasknova_test")
	os.RemoveAll(s.tempDir)
}

func (s *PermissionsTestSuite) TearDownTest() {
	os.RemoveAll(s.tempDir)
}

func (s *PermissionsTestSuite) TestEnsureStorageDirectory() {
	testPath := filepath.Join(s.tempDir, "test", "storage.json")

	err := EnsureStorageDirectory(testPath)

	assert.NoError(s.T(), err)
	info, err := os.Stat(filepath.Dir(testPath))
	assert.NoError(s.T(), err)
	assert.True(s.T(), info.IsDir())
	assert.Equal(s.T(), os.FileMode(DefaultDirMode), info.Mode().Perm())
}

func (s *PermissionsTestSuite) TestEnsureStorageDirectoryInvalidPath() {
	testPath := filepath.Join("/invalid", "path", "storage.json")

	err := EnsureStorageDirectory(testPath)

	assert.Error(s.T(), err)
}

func (s *PermissionsTestSuite) TestEnsureFilePermissions() {
	testFile := filepath.Join(s.tempDir, "test.json")
	err := os.MkdirAll(s.tempDir, DefaultDirMode)
	assert.NoError(s.T(), err)
	err = os.WriteFile(testFile, []byte("test"), 0644)
	assert.NoError(s.T(), err)

	err = EnsureFilePermissions(testFile)

	assert.NoError(s.T(), err)
	info, err := os.Stat(testFile)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), os.FileMode(DefaultFileMode), info.Mode().Perm())
}

func (s *PermissionsTestSuite) TestEnsureFilePermissionsNonExistentFile() {
	testFile := filepath.Join(s.tempDir, "nonexistent.json")

	err := EnsureFilePermissions(testFile)

	assert.Error(s.T(), err)
}

func TestPermissionsSuite(t *testing.T) {
	suite.Run(t, new(PermissionsTestSuite))
}
