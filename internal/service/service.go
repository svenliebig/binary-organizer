package service

import (
	"io/fs"
	"os"
	"path"

	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/config"
	"github.com/svenliebig/seq"
)

type service struct {
	binary binaries.Binary
	config config.Config
}

func New(binary binaries.Binary) (*service, error) {
	config, err := config.Load()
	return &service{
		binary: binary,
		config: config,
	}, err
}

// returns the installed versions of the binary that was passed to the service.
//
// possible errors:
//   - boo.ErrBinaryDirNotExists
//   - boo.ErrBinaryDirIsFile
func (s *service) Versions() ([]binaries.Version, error) {
	p, err := s.getBinaryDir()

	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(p)

	if err != nil {
		// TODO log
		return nil, err
	}

	return seq.Reduce(
		seq.Filter(
			seq.From(entries),
			func(e fs.DirEntry) (bool, error) {
				return e.IsDir(), nil
			},
		),
		func(acc []binaries.Version, e fs.DirEntry) ([]binaries.Version, error) {
			if acc == nil {
				acc = make([]binaries.Version, 0)
			}

			v, ok := s.binary.Matches(e.Name())

			if ok {
				acc = append(acc, v)
			}

			return acc, nil
		},
	)
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
