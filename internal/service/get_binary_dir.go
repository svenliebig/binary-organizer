package service

import (
	"os"
	"path"

	"github.com/svenliebig/binary-organizer/internal/boo"
)

// returns the path where the versions of a binary are located.
// the path will be secured and you can be sure the directory exists.
func (s *service) getBinaryDir() (string, error) {
	dir := path.Join(s.config.BinaryRoot, s.binary.Identifier())
	stat, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return "", boo.ErrBinaryDirNotExists
	}

	if !stat.IsDir() {
		return "", boo.ErrBinaryDirIsFile
	}

	return dir, nil
}
