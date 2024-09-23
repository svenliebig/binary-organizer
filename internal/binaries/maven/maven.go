package maven

import (
	"fmt"
	"path"
	"regexp"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

type binary struct {
}

func (n binary) Identifier() string {
	return "maven"
}

func init() {
	binaries.Register(&binary{})
}

func (n binary) BinPath(root string) string {
	return path.Join(root, "bin")
}

var FULL_VERSION_REGEX = `maven-(\d+)\.(\d+)\.(\d+)`

func (n *binary) Matches(p string) (binaries.Version, bool) {
	regexp := regexp.MustCompile(FULL_VERSION_REGEX)
	matches := regexp.FindStringSubmatch(p)

	if len(matches) != 0 {
		version := fmt.Sprintf("%s.%s.%s", matches[1], matches[2], matches[3])
		v, err := binaries.VersionFrom(version)

		if err != nil {
			return nil, false
		}

		return v, true
	}

	return nil, false
}
