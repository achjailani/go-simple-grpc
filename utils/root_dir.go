package utils

import (
	"path/filepath"
	"runtime"
)

// RootDir is function to get root directory
func RootDir() string {
	_, b, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Join(filepath.Dir(b), "..")
	}

	return ""
}
