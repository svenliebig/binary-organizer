package boo

import "errors"

var (
	// ErrBinaryDirNotExists is returned when the binary directory does not exist.
	//
	// Idea: This could be a struct containing the path that does not exist.
	// Idea: This could be a struct containing available actions to fix the error.
	// Idea: This could be a struct implementing the error and a directory not exists interface.
	ErrBinaryDirNotExists = errors.New("binary directory does not exist")
	// ErrBinaryDirIsFile is returned when the binary directory is a file.
	ErrBinaryDirIsFile = errors.New("binary directory is a file")
)
