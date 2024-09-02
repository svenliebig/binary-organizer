package service

import (
	"os"
	"path"
	"syscall"

	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/logging"
)

// returns the path where the versions of a binary are located.
// the path will be secured and you can be sure the directory exists.
//
// possible errors:
//   - boo.ErrBinaryDirNotExists
//   - boo.ErrBinaryDirIsFile
func (s *service) getBinaryDir() (string, error) {
	defer logging.Fn("service.getBinaryDir")()

	dir := path.Join(s.config.BinaryRoot, s.binary.Identifier())
	stat, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			return "", boo.ErrBinaryDirNotExists
		}
		
		if err.(*os.PathError).Err.Error() == syscall.ENOTDIR.Error() {
			return "", boo.ErrBinaryDirIsFile
		}

		logging.Error("could not get the stat of the binary directory", err)
		return "", err
	}

	if !stat.IsDir() {
		return "", boo.ErrBinaryDirIsFile
	}

	return dir, nil
}
