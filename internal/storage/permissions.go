package storage

import (
	"os"
	"path/filepath"
)

const (
	DefaultFileMode = 0600
	DefaultDirMode  = 0700
)

func EnsureStorageDirectory(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, DefaultDirMode)
}

func EnsureFilePermissions(path string) error {
	return os.Chmod(path, DefaultFileMode)
}
