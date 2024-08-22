package service

import (
	"path"
	"runtime"
)

func getValidBinaryTestdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "..", "..", "testdata", "binaries")
}

func getInvalidBinaryTestdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "..", "..", "testdata", "invalid")
}

func getInvalidBinaryTestdataFile() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "..", "..", "testdata", "error_file")
}
