package service

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

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
	defer logging.Fn("service.New")()

	c, err := config.Load()

	if err != nil {
		if errors.Is(err, boo.ErrConfigFileNotExists) {
			c, err = config.Create()
		}
	}

	if err != nil {
		logging.Error("could not load configuration", err)
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
//
// possible errors:
//   - boo.ErrVersionNotInstalled
//   - boo.ErrBinaryDirNotExists
//   - boo.ErrBinaryDirIsFile
func (s *service) getBinPath(v binaries.Version) (string, error) {
	defer logging.Fn("service.getBinPath")()

	p, err := s.getBinaryDir()

	if err != nil {
		return "", err
	}

	logging.Infof("binary directory %q", p)
	entries, err := os.ReadDir(p)

	if err != nil {
		logging.Error("could not read the binary directory", err)
		return "", err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if vers, ok := s.binary.Matches(entry.Name()); ok {
			if v.Matches(vers) {
				return s.binary.BinPath(path.Join(p, entry.Name())), nil
			}
		}
	}

	return "", boo.ErrVersionNotInstalled
}

// returns the default version of the binary that was passed to the service.
// the default version is read from the configuration file.
//
// possible errors:
//   - boo.ErrNoDefaultVersion
func (s *service) GetDefaultVersion() (binaries.Version, error) {
	defer logging.Fn("service.GetDefaultVersion")()

	versionstr, err := s.config.DefaultVersion(s.binary.Identifier())
	logging.Infof("default version %q for %q", versionstr, s.binary.Identifier())

	if err != nil {
		return nil, err
	}

	v, err := binaries.VersionFrom(versionstr)
	logging.Infof("version %q", v)

	if err != nil {
		logging.Error("could not create version from string", err)
		return nil, err
	}

	return v, nil
}

// adds a version of the configured binary to the given shell.Path parameter and removes
// all other versions of the binary from the path.
//
// first it will check if the version is instelled, if not, it will
// return a boo.ErrVersionNotInstalled error.
//
// possible errors:
//   - boo.ErrVersionNotInstalled
//   - boo.ErrBinaryDirNotExists
//   - boo.ErrBinaryDirIsFile
func (s *service) SetVersion(version binaries.Version, p shell.Path) error {
	defer logging.Fn("service.SetVersion")()

	binp, err := s.getBinPath(version)

	if err != nil {
		return err
	}

	// remove versions from path
	for _, pth := range p.Find(func(p string) bool {
		_, ok := s.binary.Matches(p)
		return ok
	}) {
		p.Remove(pth)
	}

	p.Add(binp)

	if s.binary.Identifier() == "java" {
		var home string
		if strings.HasSuffix(binp, "Contents/Home/bin") {
			home = filepath.Join(binp, "..", "..", "..")
		} else {
			home = filepath.Join(binp, "..")
		}

		return shell.WriteEnv("JAVA_HOME", home)
	}

	return nil
}
