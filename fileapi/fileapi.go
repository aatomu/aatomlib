package fileapi

import (
	"path/filepath"
	"runtime"
)

// Get Current Dir?
func CurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}
