package service

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/config"
	"github.com/svenliebig/binary-organizer/internal/logging"
	"github.com/svenliebig/binary-organizer/internal/shell"
	"github.com/svenliebig/seq"
)

type service struct {
	binary binaries.Binary
	config *config.Config
}

func New(binary binaries.Binary) (*service, error) {
	c, err := config.Load()

	if err != nil {
		if errors.Is(err, boo.ErrConfigFileNotExists) {
			c, err = config.Create()
		}
	}

	if err != nil {
		return nil, err
	}

	return &service{
		binary: binary,
		config: c,
	}, err
}

// returns the installed versions of the binary that was passed to the service.
//
// possible errors:
//   - boo.ErrBinaryDirNotExists
//   - boo.ErrBinaryDirIsFile
func (s *service) VersionsSeq() seq.Seq[binaries.Version] {
	p, err := s.getBinaryDir()

	if err != nil {
		return seq.Error[binaries.Version](err)
	}

	entries, err := os.ReadDir(p)

	if err != nil {
		logging.Error("could not read the binary directory", err)
		return seq.Error[binaries.Version](errors.New("could not read binary directory"))
	}

	return seq.FilterMap(
		seq.From(entries),
		func(e fs.DirEntry) (bool, binaries.Version, error) {
			if !e.IsDir() {
				return false, nil, nil
			}

			v, ok := s.binary.Matches(e.Name())
			return ok, v, nil
		},
	)
}

func (s *service) Versions() ([]binaries.Version, error) {
	return seq.Collect(s.VersionsSeq())
}

// checks if the binary is installed and returns the path to the binary directory.
func (s *service) IsInstalled(v binaries.Version) (string, bool) {
	p, err := s.getBinaryDir()

	if err != nil {
		// TODO log
		return "", false
	}

	entries, err := os.ReadDir(p)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if vers, ok := s.binary.Matches(entry.Name()); ok {
			if v.Matches(vers) {
				return s.binary.BinPath(path.Join(p, entry.Name())), true
			}
		}
	}

	return "", false
}

// sets a version of the configured binary in the PATH variable.
//
// first it will check if the version is instelled, if not, it will
// return a boo.ErrVersionNotInstalled error.
func (s *service) SetVersion(version binaries.Version) error {
	binp, ok := s.IsInstalled(version)

	if !ok {
		return boo.ErrVersionNotInstalled
	}

	p := shell.NewPath()

	// remove versions from path
	for _, pth := range p.Find(func(p string) bool {
		_, ok := s.binary.Matches(p)
		return ok
	}) {
		p.Remove(pth)
	}

	p.Add(binp)

	return shell.WritePath(p)
}
