package node

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

var operatingSystems = []string{"darwin", "win", "linux"}
var architectures = map[string][]string{
	"darwin": {"arm64", "x64"},
}

var FULL_VERSION_REGEX = fmt.Sprintf(`node-v(\d+)\.(\d+)\.(\d+)-(%s)-([\w\d]*)`, strings.Join(operatingSystems, "|"))

func (n *binary) Matches(p string) (binaries.Version, bool) {
	regexp := regexp.MustCompile(FULL_VERSION_REGEX)
	matches := regexp.FindStringSubmatch(p)

	if len(matches) != 0 {
		version := fmt.Sprintf("%s.%s.%s", matches[1], matches[2], matches[3])
		opers := matches[4]
		arch := matches[5]

		fmt.Println(version, opers, arch)

		v, err := binaries.VersionFrom(version)

		if err != nil {
			return nil, false
		}

		return v, true
	}

	return nil, false
}
