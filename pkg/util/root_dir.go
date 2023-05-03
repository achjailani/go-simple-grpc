package util

import (
	"os"
	"path"
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

// NewRootDir is a function
func NewRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	return dir
}
