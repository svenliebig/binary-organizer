package service

import (
	"errors"
	"path"
	"strings"
	"testing"

	"github.com/svenliebig/binary-organizer/internal/binaries"
	_ "github.com/svenliebig/binary-organizer/internal/binaries/node"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/config"
)

func TestGetBinaryDir(t *testing.T) {
	t.Run("should return the directory path when the directory exists", func(t *testing.T) {
		node, err := binaries.Get("node")

		if err != nil {
			t.Fatalf("could not get binary: %v", err)
		}

		s := service{
			binary: node,
			config: config.Config{
				BinaryRoot: getValidBinaryTestdataDir(),
			},
		}

		p, err := s.getBinaryDir()

		if err != nil {
			t.Fatalf("could not get binary directory: %v, the following directory has been checked %q", err, s.config.BinaryRoot)
		}

		suffix := path.Join("testdata", "binaries", "node")
		if !strings.HasSuffix(p, suffix) {
			t.Fatalf("directory path does not end with %v: %v", suffix, p)
		}
	})

	t.Run("should return an ErrBinaryDirNotExists when the directory does not exist", func(t *testing.T) {
		node, err := binaries.Get("node")

		if err != nil {
			t.Fatalf("could not get binary: %v", err)
		}

		s := service{
			binary: node,
			config: config.Config{
				BinaryRoot: getInvalidBinaryTestdataDir(),
			},
		}

		_, err = s.getBinaryDir()

		if err == nil {
			t.Fatalf("expected an error, but got nil")
		}

		if !errors.Is(err, boo.ErrBinaryDirNotExists) {
			t.Fatalf("expected error %v, but got %v", boo.ErrBinaryDirNotExists, err)
		}
	})

	t.Run("should return an ErrBinaryDirIsFile when the directory is a file", func(t *testing.T) {
		node, err := binaries.Get("node")

		if err != nil {
			t.Fatalf("could not get binary: %v", err)
		}

		s := service{
			binary: node,
			config: config.Config{
				BinaryRoot: getInvalidBinaryTestdataFile(),
			},
		}

		_, err = s.getBinaryDir()

		if err == nil {
			t.Fatalf("expected an error, but got nil")
		}

		if !errors.Is(err, boo.ErrBinaryDirIsFile) {
			t.Fatalf("expected error %v, but got %v", boo.ErrBinaryDirIsFile, err)
		}
	})
}
