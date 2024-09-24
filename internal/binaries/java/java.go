package java

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

type binary struct {
}

func (n binary) Identifier() string {
	return "java"
}

func init() {
	binaries.Register(&binary{})
}

func (n binary) BinPath(root string) string {
	if _, err := os.Stat(path.Join(root, "Contents", "Home")); err == nil {
		return path.Join(root, "Contents", "Home", "bin")
	}

	return path.Join(root, "bin")
}

var _ = `/jdk-?(\d+)(?:\.(\d+))?(?:\.(\d+))?(?:_(\d+))?(?:[^/]*)/.*`
var FULL_VERSION_REGEX = `jdk-?(\d+)(?:\.(\d+))?(?:\.(\d+))?`

func (n *binary) Matches(p string) (binaries.Version, bool) {
	regexp := regexp.MustCompile(FULL_VERSION_REGEX)
	matches := regexp.FindStringSubmatch(p)

	if len(matches) != 0 {
		major := getIndex(matches, 1)
		minor := getIndex(matches, 2)
		fix := getIndex(matches, 3)

		if major == "1" {
			major = minor
			minor = "0"
		}

		version := fmt.Sprintf("%s.%s.%s", major, minor, fix)
		v, err := binaries.VersionFrom(version)

		if err != nil {
			return nil, false
		}

		return v, true
	}

	return nil, false
}

func getIndex(m []string, index int) string {
	if len(m) > index {
		r := strings.Trim(m[index], " ")

		if r != "" {
			return r
		}
	}

	return "0"
}
